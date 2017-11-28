# 程序控制结构

虽然剧透可耻，但是为了体现 Go 语言的设计简洁之处，必须要先剧透一下。

Go 语言的控制结构关键字只有

`if..else if..else`，`for` 和 `switch`。

而且在 Go 中，为了避免格式化战争，对程序结构做了统一的强制的规定。看下下面的例子。

请比较一下 A 程序和 B 程序的不同之处。

**A 程序**

    package main

    import (
    	"fmt"
    )

    func main() {
    	fmt.Println("hello world")
    }

**B 程序**

    package main

    import (
    	"fmt"
    )

    func main()
    {
    	fmt.Println("hello world")
    }

还记得我们前面的例子中，`{}`的格式是怎么样的么？在上面的两个例子中只有 A 例的写法是对的。因为在 Go 语言中，强制了`{}`的格式。如果我们试图去编译 B 程序，那么会发生如下的错误提示。

    ./test_format.go:9: syntax error: unexpected semicolon or newline before {

**if..else if..else**

if..else if..else 用来判断一个或者多个条件，然后根据条件的结果执行不同的程序块。举个简单的例子。

    package main

    import (
    	"fmt"
    )

    func main() {
    	var dog_age = 10

    	if dog_age > 10 {
    		fmt.Println("A big dog")
    	} else if dog_age > 1 && dog_age <= 10 {
    		fmt.Println("A small dog")
    	} else {
    		fmt.Println("A baby dog")
    	}
    }

上面的例子判断狗狗的年龄如果`(if)`大于 10 就是一个大狗；否则判断`(else if)`狗狗的年龄是否小于等于 10 且大于 1，这个时候狗狗是小狗狗。否则`(else)`的话（就是默认狗狗的年龄小于等于 1 岁），那么狗狗是 Baby 狗狗。

在上面的例子中，我们还可以发现 Go 的 if..else if..else 语句的判断条件一般都不需要使用`()`。当然如果你还是愿意写，也是对的。另外如果为了将某两个或多个条件绑定在一起判断的话，还是需要括号`()`的。

比如下面的例子也是对的。

    package main

    import (
    	"fmt"
    )

    func main() {
    	const Male = 'M'
    	const Female = 'F'

    	var dog_age = 10
    	var dog_sex = 'M'

    	if (dog_age == 10 && dog_sex == 'M') {
    		fmt.Println("dog")
    	}
    }

但是如果你使用 Go 提供的格式化工具来格式化这段代码的话，Go 会智能判断你的括号是否必须有，否则的话，会帮你去掉的。你可以试试。

    go fmt test_bracket.go

然后你会发现，咦？！果真被去掉了。

另外因为每个判断条件的结果要么是 true 要么是 false，所以可以使用`&&`，`||`来连接不同的条件。使用`!`来对一个条件取反。

**switch**

switch 的出现是为了解决某些情况下使用 if 判断语句带来的繁琐之处。

例如下面的例子：

    package main

    import (
    	"fmt"
    )

    func main() {
    	//score 为 [0,100]之间的整数
    	var score int = 69

    	if score >= 90 && score <= 100 {
    		fmt.Println("优秀")
    	} else if score >= 80 && score < 90 {
    		fmt.Println("良好")
    	} else if score >= 70 && score < 80 {
    		fmt.Println("一般")
    	} else if score >= 60 && score < 70 {
    		fmt.Println("及格")
    	} else {
    		fmt.Println("不及格")
    	}
    }

在上面的例子中，我们用 if..else if..else 来对分数进行分类。这个只是一般的情况下 if 判断条件的数量。如果 if..else if..else 的条件太多的话，我们可以使用 switch 来优化程序。比如上面的程序我们还可以这样写：

    package main

    import (
    	"fmt"
    )

    func main() {
    	//score 为 [0,100]之间的整数
    	var score int = 69

    	switch score / 10 {
    	case 10:
    	case 9:
    		fmt.Println("优秀")
    	case 8:
    		fmt.Println("良好")
    	case 7:
    		fmt.Println("一般")
    	case 6:
    		fmt.Println("及格")
    	default:
    		fmt.Println("不及格")
    	}
    }

关于 switch 的几点说明如下：

(1) switch 的判断条件可以为任何数据类型。

    package main

    import (
    	"fmt"
    )

    func main() {
    	var dog_sex = "F"
    	switch dog_sex {
    	case "M":
    		fmt.Println("A male dog")
    	case "F":
    		fmt.Println("A female dog")
    	}
    }

(2) 每个 `case` 后面跟的是一个完整的程序块，该程序块`不需要 {}`，也`不需要 break 结尾`，因为每个 `case` 都是独立的。

(3) 可以为 `switch` 提供一个默认选项 `default`，在上面所有的 `case` 都没有满足的情况下，默认执行 `default` 后面的语句。

**for**

for 用在 Go 语言的循环条件里面。比如说要你输出 1...100 之间的自然数。最笨的方法就是直接这样。

    package main

    import (
    	"fmt"
    )

    func main() {
    	fmt.Println(1)
    	fmt.Println(2)
    	...
    	fmt.Println(100)
    }

这个不由地让我想起一个笑话。

> 以前一个地主的儿子学习写字，只学了三天就把老师赶走了。因为在这三天里面他学写了一，二，三。他觉得写字真的太简单了，不就是画横线嘛。于是有一天老爹过寿，让他来记送礼的人名单。直到中午还没有记完，老爹很奇怪就去问他怎么了。他哭着说，“ 不知道这个人有什么毛病，姓什么不好，姓万 ”。

哈哈，回来继续。我们看到上面的例子也是如地主的儿子那样就不好了。所以，我们必须使用循环结构。我们用 for 的循环语句来实现上面的例子。

    package main

    import (
    	"fmt"
    )

    func main() {
    	var i int = 1

    	for ; i <= 100; i++ {
    		fmt.Println(i)
    	}
    }

在上面的例子中，首先初始化变量 i 为 1，然后在 for 循环里面判断是否小于等于 100，如果是的话，输出 i，然后再使用 i++ 来将 i 的值自增 1。上面的例子，还有一个更好的写法，就是将 i 的定义和初始化也放在 for 里面。如下：

    package main

    import (
    	"fmt"
    )

    func main() {
    	for i := 1; i <= 100; i++ {
    		fmt.Println(i)
    	}
    }

在 Go 里面没有提供 while 关键字，如果你怀念 while 的写法也可以这样：

    package main

    import (
    	"fmt"
    )

    func main() {
    	var i int = 1

    	for i <= 100 {
    		fmt.Println(i)
    		i++
    	}
    }

或许你会问，如果我要死循环呢？是不是`for true`？呵呵，不用了，直接这样。

    for{
    	...
    }

以上就是 Go 提供的全部控制流程了。

再复习一下，Go 只提供了：

**if**

    if ...{
    	...
    }else if ...{
    	...
    }else{
    	...
    }

**switch** switch(...){ case ...: ... case ...: ... ... default: ... } **for** for ...; ...; ...{ ... } for ...{ ... } for{ ... }
