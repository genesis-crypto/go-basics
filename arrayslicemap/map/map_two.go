package main

import "fmt"

func main() {
	var salaries map[string]float64 = map[string]float64{
		"lex": 23001.23,
		"ska": 1239.21,
	}

	salaries["ush"] = 2193.52

	delete(salaries, "lex")

	fmt.Println(salaries)

	for name, salary := range salaries {
		fmt.Printf("Name: %v - Salary: %v\n", name, salary)
	}
}
