package main

import (
	"fmt"
)

var x int
var y float64

var a int8 = -128

func main() {
	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("%T\n", x)
	fmt.Println("%T\n", y)

}
