package main

import "fmt"

func main() {
	i := 0

	for i <= 10 {
		fmt.Printf("%v", i)
		i++
	}

	for j := 10; j <= 20; j++ {
		fmt.Printf("%v", j)
	}

	for {
		fmt.Println("Loooop")
	}
}
