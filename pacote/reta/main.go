package main

import "fmt"

func main() {
	p1 := Point{2.0, 2.0}
	p2 := Point{2.0, 2.0}

	fmt.Println(sides(p1, p2))
	fmt.Println(p1, p2)
}
