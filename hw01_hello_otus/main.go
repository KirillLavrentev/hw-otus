package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	var hello string = "Hello, OTUS!"
	fmt.Println(reverse.String(hello))

}
