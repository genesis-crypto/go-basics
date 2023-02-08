package main

import "fmt"

func gradeByRank(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 7 && grade < 9 {
		return "B"
	} else if grade >= 5 && grade < 7 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	fmt.Println(gradeByRank(10))
	fmt.Println(gradeByRank(7))
	fmt.Println(gradeByRank(5))
	fmt.Println(gradeByRank(2))
}
