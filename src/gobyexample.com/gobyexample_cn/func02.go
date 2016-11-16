/*  project: gobyexample_cn
    author:diogo
    time: 2016/11/16-13:52
*/
/**
Go 闭包函数
Go支持匿名函数，匿名函数可以形成闭包。闭包函数可以访问定义闭包的函数定义的内部变量。
*/

/**
示例 02
package main

import "fmt"

func main() {
	add10 := closure(10)  //其实是构造了一个加10函数
	fmt.Println(add10(5)) //15
	fmt.Println(add10(6)) //16

	add20 := closure(20)
	fmt.Println(add20(5)) //25
}

func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}

}
*/
/**
package main

import "fmt"

func main() {

	var fs []func() int

	for i := 0; i < 3; i++ {

		fs = append(fs, func() int {

			return i
		})
	}
	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}

	fmt.Println(fs)
}
*/

//输出结果
//0x4013f0 = 3
//0x4013f0 = 3
//0x4013f0 = 3
//[0x4013f0 0x4013f0 0x4013f0]


package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    result := adder()
    for i := 0; i < 5; i++ {
        fmt.Println(result(i))
    }
}
/** 输出
0
1
3
6
10
15
21
28
36
45
 */