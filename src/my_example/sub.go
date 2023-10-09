package main

import "fmt"

func main() {

	finds("userToken", "okay")

}

func finds(data ...interface{}) error {

	for key, value := range data {

		fmt.Println(key)
		fmt.Println(value)

	}

	return nil
}
