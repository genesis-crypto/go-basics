package main

import "fmt"

func first_func() {
	fmt.Println("First")
}

func second_func(param_one string, param_two string) {
	fmt.Printf("Second function: %s %s\n", param_one, param_two)
}

func third_func() string {
	return "third function"
}

func fourth_func(param_one, param_two string) string {
	return fmt.Sprintf("Fourth: %s %s", param_one, param_two)
}

func fifty_func() (string, string) {
	return "Return one", "Return two"
}

func main() {
	first_func()
	second_func("a", "b")

	third_result, fourth_result := third_func(), fourth_func("a", "b")
	fmt.Println(third_result, fourth_result)

	returnOne, returnTwo := fifty_func()
	fmt.Println("Fifty: ", returnOne, returnTwo)
}
