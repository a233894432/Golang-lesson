package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("%q\n", strings.Split("foo|bar|baz", "|"))

	var f interface{}
	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Morticia",
			"Gomez",
		},
	}

	data := f.(map[string]interface{})

	fmt.Println(data)

}
