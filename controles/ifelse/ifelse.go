package main

import "fmt"

func showResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved <> ", grade)
		return
	}
	fmt.Println("Rejected <> ", grade)
}

func main() {
	showResult(8)
	showResult(6)
}
