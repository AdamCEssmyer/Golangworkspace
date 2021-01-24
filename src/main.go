package main

import (
	"fmt"
)

var a int

type hamburger int

var b hamburger

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Println(b)
	a = int(b)
}
