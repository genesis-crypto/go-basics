package main

import "fmt"

func gradeByRank(grade float64) string {
	switch {
	case grade <= 10 && grade >= 9:
		return "A"
	case grade < 9 && grade >= 8:
		return "B"
	case grade < 8 && grade >= 7:
		return "C"
	default:
		return "D"
	}
}

func main() {
	fmt.Println(gradeByRank(10))
	fmt.Println(gradeByRank(20))
}
