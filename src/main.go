package main

import "fmt"

var y = 43

func main() {
	//short declaration operator
	//Declare a variable and Assign a Value
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("again:", y)
}
