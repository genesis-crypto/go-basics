package main

import (
	"fmt"
	"reflect"
)

func main() {
	array_numbers := [3]int{1, 2, 3}
	slice_numbers := []int{1, 2, 3}

	fmt.Println(reflect.TypeOf(array_numbers))
	fmt.Println(reflect.TypeOf(slice_numbers))

	second_array_numbers := [5]int{1, 2, 3, 4, 5}
	second_slice_number := second_array_numbers[1:3]
	second_array_numbers[1] = 9999

	fmt.Println(second_array_numbers, second_slice_number)

	third_slice_number := second_array_numbers[:2]
	fourth_slice_number := second_slice_number[:1]
	second_slice_number[0] = 100000
	fmt.Println(third_slice_number)
	fmt.Println(fourth_slice_number)
}
