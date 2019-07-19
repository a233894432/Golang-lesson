package main

import "fmt"

func main() {

	var a int
	a = 10
	judgeType(a)

	data := `{"name":"diogoxiang"}`

	judgeType(data)

}

func judgeType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
