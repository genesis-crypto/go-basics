package main

import "fmt"

func recursion(x uint) (uint, error) {
	if x == 0 {
		return 0, nil
	}

	return recursion(x - 1)
}

func main() {
	fmt.Println(recursion(10))
}
