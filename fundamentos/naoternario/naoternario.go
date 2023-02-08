package main

import "fmt"

func resultGrade(grade float64) string {
	if grade >= 6 {
		return "Approved"
	}

	return "Rejected"
}

func main() {
	fmt.Println(resultGrade(9))
}
