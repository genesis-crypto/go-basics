package main

import "fmt"

func main() {
	slice_integer := make([]int, 10)
	fmt.Println(slice_integer, len(slice_integer), cap(slice_integer))

	second_slice_integer := make([]int, 10, 20)
	fmt.Println(second_slice_integer, len(second_slice_integer), cap(second_slice_integer))

	slice_integer = append(slice_integer, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice_integer, len(slice_integer), cap(slice_integer))
}
