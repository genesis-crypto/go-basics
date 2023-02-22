package main

import "fmt"

type Grade float64

func (g Grade) between(a, b float64) bool {
	return float64(g) >= a && float64(g) <= b
}

func gradeByRank(grade Grade) string {
	switch {
	case grade.between(9.0, 10):
		return "A"
	case grade.between(7.0, 8.99):
		return "B"
	case grade.between(5.0, 7.99):
		return "C"
	default:
		return "D"
	}
}

func main() {
	fmt.Println(gradeByRank(8.0))
}
