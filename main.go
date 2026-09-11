package main

import (
	"fmt"
	"module_provider/foo"
	"module_provider/bar"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("hello %s\n", foo.Fooit())
	fmt.Printf("hello %s\n", bar.Barit())
}
