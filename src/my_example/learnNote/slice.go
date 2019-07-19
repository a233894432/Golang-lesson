package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	modify(slice)
	fmt.Println(slice)
	fmt.Println(len(slice))
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}
