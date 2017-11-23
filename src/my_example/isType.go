package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = `{"from":"en","to":"zh","trans_result":{"src":"today","dst":"\u4eca\u5929"},"result":["src","today","dst","\u4eca\u5929"]}`
	fmt.Println("type:", reflect.TypeOf(x))
}
