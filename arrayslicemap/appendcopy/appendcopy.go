package main

import "fmt"

func main() {
	array_one := [3]int{1, 2, 3}
	var slice_one []int

	slice_one = append(slice_one, 4, 5, 7)
	fmt.Printf("Array: %v <> Slice: %v\n", array_one, slice_one)

	slice_three := make([]int, 2)
	copy(slice_three, slice_one)

	fmt.Printf("Slice Three = %v", slice_three)
}
