这篇文章想聊聊 Golang 语言下的设计模式问题，我觉得这个话题还是比较有意思的。Golang 没有像 java 那样对设计模式疯狂的迷恋，而是摆出了一份 “ 看庭前花开花落，望天空云卷云舒 ” 的姿态。

## 单例模式 :

Gloang 的单例模式该怎么写？随手写一个，不错，立马写出来了。但这个代码有什么问题呢？多个协程同时执行这段代码就会出现问题：`instance`可能会被赋值多次，这段代码是线程不安全的代码。那么如何保证在多线程下只执行一次呢？条件反射：加锁。。。加锁是可以解决问题。但不是最优的方案，因为如果有 1W 并发，每一个线程都竞争锁，同一时刻只有一个线程能拿到锁，其他的全部阻塞等待。让原本想并发得飞起来变成了一切认怂串行化。通过`check-lock-check`方式可以减少竞争。还有其他方式，利用`sync/atomic`和`sync/once` 这里只给出代码

```go
func NewSingleton() *singleton {
    if instance == nil {
         instance = &singleton{}
    }
    return instance
}````
```

```go
func NewSingleton() *singleton {
    l.Lock()                   // lock
    defer l.Unlock()
    if instance == nil {  // check
        instance = &singleton{}
    }
    return instance
}
```

```go
func NewSingleton() *singleton {
    if instance == nil {    // check
        l.Lock()            // lock
        defer l.Unlock()
        if instance == nil {    // check
            instance = &singleton{}
        }
    }
    return instance
}
```

```go
func NewSingleton() *singleton {
    if atomic.LoadUInt32(&initialized) == 1 {
        return instance
    }
    mu.Lock()
    defer mu.Unlock()
    if initialized == 0 {
        instance = &singleton{}
        atomic.StoreUint32(&initialized, 1)
    }
    return instance
}
```

```go
func NewSingleton() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

## 工厂模式 :

工厂根据条件产生不同功能的类。工厂模式使用经常使用在替代`new`的场景中，让工厂统一根据不同条件生产不同的类。工厂模式在解耦方面将使用者和产品之间的依赖推给了工厂，让工厂承担这种依赖关系。工厂模式又分为简单工厂，抽象工厂。golang 实现一个简单工厂模式如下 :

```go
package main
import (
    "fmt"
)
type Op interface {
    getName() string
}
type A struct {
}
type B struct {
}
type Factory struct {
}
func (a *A) getName() string {
    return "A"
}
func (b *B) getName() string {
    return "B"
}
func (f *Factory) create(name string) Op {
    switch name {
    case `a`:
        return new(A)
    case `b`:
        return new(B)
    default:
        panic(`name not exists`)
    }
    return nil
}
func main() {
    var f = new(Factory)
    p := f.create(`a`)
    fmt.Println(p.getName())
    p = f.create(`b`)
    fmt.Println(p.getName())
}
```

## 依赖注入 :

具体含义是 : 当某个角色 ( 可能是一个实例，调用者 ) 需要另一个角色 ( 另一个实例，被调用者 ) 的协助时，在传统的程序设计过程中，通常由调用者来创建被调用者的实例。但在这种场景下，创建被调用者实例的工作通常由容器 (IoC) 来完成，然后注入调用者，因此也称为依赖注入。 Golang 利用函数 f 可以当做参数来传递，同时配合`reflect 包`拿到参数的类型，然后根据调用者传来的参数和类型匹配上之后，最后通过`reflect.Call()`执行具体的函数。下面的代码来自：https://www.studygolang.com/articles/4957 这篇文章上。

```go
package main

import (
    "fmt"
    "reflect"
)

var inj *Injector

type Injector struct {
    mappers map[reflect.Type]reflect.Value // 根据类型map实际的值
}

func (inj *Injector) SetMap(value interface{}) {
    inj.mappers[reflect.TypeOf(value)] = reflect.ValueOf(value)
}

func (inj *Injector) Get(t reflect.Type) reflect.Value {
    return inj.mappers[t]
}

func (inj *Injector) Invoke(i interface{}) interface{} {
    t := reflect.TypeOf(i)
    if t.Kind() != reflect.Func {
        panic("Should invoke a function!")
    }
    inValues := make([]reflect.Value, t.NumIn())
    for k := 0; k < t.NumIn(); k++ {
        inValues[k] = inj.Get(t.In(k))
    }
    ret := reflect.ValueOf(i).Call(inValues)
    return ret
}

func Host(name string, f func(a int, b string) string) {
    fmt.Println("Enter Host:", name)
    fmt.Println(inj.Invoke(f))
    fmt.Println("Exit Host:", name)
}

func Dependency(a int, b string) string {
    fmt.Println("Dependency: ", a, b)
    return `injection function exec finished ...`
}

func main() {
    // 创建注入器
    inj = &Injector{make(map[reflect.Type]reflect.Value)}
    inj.SetMap(3030)
    inj.SetMap("zdd")

    d := Dependency
    Host("zddhub", d)

    inj.SetMap(8080)
    inj.SetMap("www.zddhub.com")
    Host("website", d)
}
```

## 装饰器模式 :

装饰器模式：允许向一个现有的对象添加新的功能，同时又不改变其结构。这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。我们使用最为频繁的场景就是`http`请求的处理：对`http`请求做`cookie`校验。

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func autoAuth(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("Auth")
        if err != nil || cookie.Value != "Authentic" {
            w.WriteHeader(http.StatusForbidden)
            return
        }
        h(w, r)
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func main() {
    http.HandleFunc("/hello", autoAuth(hello))
    err := http.ListenAndServe(":5666", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

还有很多其他模式，这里不一一给出了，写这篇文章的目的是想看看这些模式在 `golang` 中是如何体现出来的，框架或者类库应该是设计模式常常出没的地方。深入理解设计模式有助于代码的抽象，复用和解耦，让代码与代码之间更加低耦合。
