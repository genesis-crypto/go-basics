package main

import "fmt"

func main() {
	employeesByLetters := map[string]map[string]float64{
		"A": {
			"alice": 12931.20,
			"ashi":  12993.19,
		},
		"B": {
			"bog": 2391.29,
			"bia": 2029.29,
		},
	}

	for group, employees := range employeesByLetters {
		fmt.Printf("Grupo: %v => Employees: %v\n", group, employees)
	}
}
