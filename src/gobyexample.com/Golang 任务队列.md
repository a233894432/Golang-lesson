Golang 在异步处理上有着上佳的表现。因为 goroutines 和 channels 是非常容易使用且有效的异步处理手段。下面我们一起来看一看 Golang 的简易任务队列

## 一种 " 非任务队列 " 的任务队列

有些时候，我们需要做异步处理但是并不需要一个任务对列，这类问题我们使用 Golang 可以非常简单的实现。如下：

```go
go process(job)
```

这的确是很多场景下的绝佳选择，比如操作一个 HTTP 请求等待结果。然而，在一些相对复杂高并发的场景下，你就不能简单的使用该方法来实现异步处理。这时候，你需要一个队列来管理需要处理的任务，并且按照一定的顺序来处理这些任务。

## 最简单的任务队列

接下来看一个最简单的任务队列和工作者模型。

```go
func worker(jobChan <-chan Job) {
    for job := range jobChan {
        process(job)
    }
}

// make a channel with a capacity of 100.
jobChan := make(chan Job, 100)

// start the worker
go worker(jobChan)

// enqueue a job
jobChan <- job
```

代码中创建了一个 Job 对象的 channel , 容量为 100。然后开启一个工作者协程从 channel 中去除任务并执行。任务的入队操作就是将一个 Job 对象放入任务 channel 中。

虽然上面只有短短的几行代码，却完成了很多的工作。我们实现了一个简易的线程安全的、支持并发的、可靠的任务队列。

## 限流

上面的例子中，我们初始化了一个容量为 100 的任务 channel。

```go
// make a channel with a capacity of 100.
jobChan := make(chan Job, 100)
```

这意味着任务的入队操作十分简单，如下：

```go
// enqueue a job
jobChan <- job
```

这样一来，当 job channel 中已经放入 100 个任务的时候，入队操作将会阻塞，直至有任务被工作者处理完成。这通常不是一个好的现象，因为我们通常不希望程序出现阻塞等待。这时候，我们通常希望有一个超时机制来告诉服务调用方，当前服务忙，稍后重试。我之前的博文 -- 我读《通过 Go 来处理每分钟达百万的数据请求》介绍过类似的限流策略。这里方法类似，就是当队列满的时候，返回 503，告诉调用方服务忙。代码如下：

```go
// TryEnqueue tries to enqueue a job to the given job channel. Returns true if
// the operation was successful, and false if enqueuing would not have been
// possible without blocking. Job is not enqueued in the latter case.
func TryEnqueue(job Job, jobChan <-chan Job) bool {
    select {
    case jobChan <- job:
        return true
    default:
        return false
    }
}
```

这样一来，我们尝试入队的时候，如果入队失败，放回一个 false ，这样我们再对这个返回值处理如下：

```go
if !TryEnqueue(job, chan) {
    http.Error(w, "max capacity reached", 503)
    return
}
```

这样就简单的实现了限流操作。当 jobChan 满的时候，程序会走到 default 返回 false ，从而告知调用方当前的服务器情况。

## 关闭工作者

到上面的步骤，限流已经可以解决，那么我们接下来考虑，怎么才能优雅的关闭工作者？假设我们决定不再向任务队列插入任务，我们希望让所有的已入队任务执行完成，我们可以非常简单的实现：

```go
close(jobChan)
```

没错，就是这一行代码，我们就可以让任务队列不再接收新任务（仍然可以从 channel 读取 job ），如果我们想执行队列里的已经存在的任务，只需要：

```go
for job := range jobChan {...}
```

所有已经入队的 job 会正常被 woker 取走执行。但是，这样实际上还存在一个问题，就是主协成不会等待工作者执行完工作就会退出。它不知道工作者协成什么时候能够处理完以上的任务。可以运行的例子如下 :

```go
package main

import (
    "fmt"
)

var jobChan chan int

func worker(jobChan <- chan int)  {
    for job := range jobChan{
        fmt.Printf("执行任务 %d \n", job)
    }
}

func main() {
    jobChan = make(chan int, 100)
    //入队
    for i := 1; i <= 10; i++{
        jobChan <- i
    }

    close(jobChan)
    go worker(jobChan)

}
```

运行发现，woker 无法保证执行完 channel 中的 job 就退出了。那我们怎么解决这个问题？

## 等待 woker 执行完成

使用 sysc.WaitGroup:

```go
package main

import (
    "fmt"
    "sync"
)

var jobChan chan int
var wg sync.WaitGroup

func worker(jobChan <- chan int)  {
    defer wg.Done()
    for job := range jobChan{
        fmt.Printf("执行任务 %d \n", job)
    }
}

func main() {
    jobChan = make(chan int, 100)
    //入队
    for i := 1; i <= 10; i++{
        jobChan <- i
    }

    wg.Add(1)
    close(jobChan)

    go worker(jobChan)
    wg.Wait()
}
```

使用这种协程间同步的方法，协成会等待 worker 执行完 job 才会退出。运行结果：

```go
执行任务 1
执行任务 2
执行任务 3
执行任务 4
执行任务 5
执行任务 6
执行任务 7
执行任务 8
执行任务 9
执行任务 10

Process finished with exit code 0
```

这样是完美的么？在设计功能的时候，为了防止协程假死，我们应该给协程设置一个超时。

## 超时设置

上面的例子中 wg.Wait() 会一直等待，直到 wg.Done() 被调用。但是如果这个操作假死，无法调用，将永远等待。这是我们不希望看到的，因此，我们可以给他设置一个超时时间。方法如下：

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var jobChan chan int
var wg sync.WaitGroup

func worker(jobChan <-chan int) {
    defer wg.Done()
    for job := range jobChan {
        fmt.Printf("执行任务 %d \n", job)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    jobChan = make(chan int, 100)
    //入队
    for i := 1; i <= 10; i++ {
        jobChan <- i
    }

    wg.Add(1)
    close(jobChan)

    go worker(jobChan)
    res := WaitTimeout(&wg, 5*time.Second)
    if res {
        fmt.Println("执行完成退出")
    } else {
        fmt.Println("执行超时退出")
    }
}

//超时机制
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
    ch := make(chan struct{})
    go func() {
        wg.Wait()
        close(ch)
    }()
    select {
    case <-ch:
        return true
    case <-time.After(timeout):
        return false
    }
}
```

```go
执行结果如下：

执行任务 1
执行任务 2
执行任务 3
执行任务 4
执行任务 5
执行超时退出

Process finished with exit code 0
```

这样，5s 超时生效，虽然不是所有的任务被执行，由于超时，也会退出。

有时候我们希望 woker 丢弃在执行的工作，也就是 cancel 操作，怎么处理？

## Cancel Worker

我们可以借助 context.Context 实现。如下：

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

var jobChan chan int
var ctx context.Context
var cancel context.CancelFunc

func worker(jobChan <-chan int, ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        case job := <-jobChan:
            fmt.Printf("执行任务 %d \n", job)
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    jobChan = make(chan int, 100)
    //带有取消功能的 contex
    ctx, cancel = context.WithCancel(context.Background())
    //入队
    for i := 1; i <= 10; i++ {
        jobChan <- i
    }

    close(jobChan)

    go worker(jobChan, ctx)
    time.Sleep(2 * time.Second)
    //調用cancel
    cancel()
}
```

```go
結果如下：

执行任务 1
执行任务 2

Process finished with exit code 0
```

可以看出，我们等待 2s 后，我们主动调用了取消操作，woker 协程主动退出。

这是借助 context 包实现了取消操作，实质上也是监听一个 channel 的操作，那我们有没有可能不借助 context 实现取消操作呢？

不使用 context 的超时机制实现取消：

```go
package main

import (
    "fmt"
    "time"
)

var jobChan chan int

func worker(jobChan <-chan int, cancelChan <-chan struct{}) {
    for {
        select {
        case <-cancelChan:
            return
        case job := <-jobChan:
            fmt.Printf("执行任务 %d \n", job)
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    jobChan = make(chan int, 100)
    //通过chan 取消操作
    cancelChan := make(chan struct{})
    //入队
    for i := 1; i <= 10; i++ {
        jobChan <- i
    }

    close(jobChan)

    go worker(jobChan, cancelChan)
    time.Sleep(2 * time.Second)
    //关闭chan
    close(cancelChan)
}
```

这样，我们使用一个关闭 chan 的信号实现了取消操作。原因是无缓冲 chan 读取会阻塞，当关闭后，可以读取到空，因此会执行 select 里的 return.

## 总结

照例总结一波，本文介绍了 golang 协程间的同步和通信的一些方法，任务队列的最简单实现。关于工作者池的实现，我在其他博文也写到了，这里不多写。本文更多是工具性的代码，写功能时候可以借用，比如超时、取消、chan 的操作等。
