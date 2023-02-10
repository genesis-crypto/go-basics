package main

import "fmt"

func main() {
	var array_number [4]float64 = [4]float64{8.1, 7.2, 9.0, 10.0}
	var first_slice_number = array_number[:4]
	var second_slice_number = array_number[:4]
	array_number[1] = 6.3

	fmt.Printf("1 - Slice %v\n", first_slice_number)
	fmt.Printf("2 - Slice %v\n", second_slice_number)

	var fourth_slice_number = make([]float64, 0, len(array_number))
	fourth_slice_number = array_number[:]

	fmt.Printf("4 - Slice %v\n", fourth_slice_number)
}
