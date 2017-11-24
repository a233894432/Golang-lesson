package main

import (
	"fmt"
)

func main() {
	fmt.Println(whatType("12"))
}

func whatType(x interface{}) string {

	switch x.(type) {
	case int:
		return "int"
	default:
		return "null"

	}

}
