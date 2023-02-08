package main

import "fmt"

func gradeByRank(_grade float64) string {
	var grade = int(_grade)
	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid Grade"
	}
}

func main() {
	fmt.Println(gradeByRank(10))
	fmt.Println(gradeByRank(11))
	fmt.Println(gradeByRank(9))
	fmt.Println(gradeByRank(8.9))
}
