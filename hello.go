package main

import "fmt"

func hello(name string) string {
	// return "Hello, " + name + "!"
	return fmt.Sprintf("Hello, %s!", name)
}
