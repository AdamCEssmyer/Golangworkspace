package main

import (
	"fmt"
)

type age int

var x age

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)
}
