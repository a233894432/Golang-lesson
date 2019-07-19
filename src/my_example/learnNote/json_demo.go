package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})
	fmt.Println(data)

	for k, v := range data {
		// fmt.Println(k, v)
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
			data[k] = RandomInt(10)
			fmt.Println(data)
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}

	fmt.Println(RandomInt(10))

}

func RandomInt(n int) int {

	return rand.Intn(n)
}
