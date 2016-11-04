package main

import "fmt"

//示例代码
var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

func main() {

	fmt.Println(isActive)

	s := "hello,"
	m := " world"
	a := s + m[1:]
	fmt.Printf("%s\n", a)

}
