package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	str := "Hello, OTUS!"
	rStr := reverse.String(str)
	fmt.Print(rStr)
	// Place your code here.
}
