package main

import "fmt"

func switcher(one, two int) (first, second int) {
	second = one
	first = two
	return
}

func main() {
	second, first := switcher(2, 3)
	fmt.Println(second, first)
}
