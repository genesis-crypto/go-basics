package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func exec(function func(a, b int) int, a, b int) int {
	return function(a, b)
}

func main() {
	fmt.Println(exec(mult, 3, 2))
}
