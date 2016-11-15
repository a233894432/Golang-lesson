/*  project: my_example
    author:diogo
    time: 2016/11/15-19:49
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{"diogoxiang", 20}

	fmt.Println(p)
	//输出 指针 内存地址
	fmt.Printf("%p\n",&p)

	// 这个语法创建一个新结构体变量
	fmt.Println(Person{"Bob", 20})

	// 可以使用"成员:值"的方式来初始化结构体变量
	fmt.Println(Person{Name: "Alice", Age: 30})

	// 未显式赋值的成员初始值为零值
	fmt.Println(Person{Name: "Fred"})

	// 可以使用&来获取结构体变量的地址
	fmt.Printf("%p \n",&Person{Name: "Ann", Age: 40})  //0xc042008740

	// 使用点号(.)来访问结构体成员
	s := Person{Name: "Sean", Age: 50}
	fmt.Println(s.Name)

	// 结构体指针也可以使用点号(.)来访问结构体成员
	// Go语言会自动识别出来
	sp := &s
	fmt.Println(sp.Age)

	// 结构体成员变量的值是可以改变的
	sp.Age = 51
	fmt.Println(sp.Age)


}
