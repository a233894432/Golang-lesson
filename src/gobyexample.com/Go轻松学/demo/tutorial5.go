package main

import (
	"fmt"
)

func main() {
	var x = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday"}

	for _, day := range x {
		fmt.Println(day)
	}
}
