package main

import "fmt"

func increment_one(n int) {
	n++
}

func increment_two(n *int) {
	*n++
}

func main() {
	x := 20
	increment_one(x)
	increment_two(&x)
	fmt.Println("x: ", x)
}
