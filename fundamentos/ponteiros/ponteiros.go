package main

import "fmt"

func main() {
	i := 1

	var p *int = &i
	*p++
	i++
	fmt.Println(*p, p, i, &i)
}
