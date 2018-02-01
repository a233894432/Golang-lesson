package main

import (
	"fmt"
)

func main() {
	fmt.Println(whatType("{id:12}"))
}

func whatType(x interface{}) string {

	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "null"

	}

}
