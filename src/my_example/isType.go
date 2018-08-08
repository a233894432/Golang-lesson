package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var x = `{"from":"en","to":"zh","trans_result":{"src":"today","dst":"\u4eca\u5929"},"result":["src","today","dst","\u4eca\u5929"]}`
	var n = 123
	fmt.Println("type:", reflect.TypeOf(n))
}
