package main

import "fmt"

func main() {
	var h string

	h = hello("World")
	fmt.Println(h)
	h1 := hello("Andrey")
	fmt.Println(h1)
}
