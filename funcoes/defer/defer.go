package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("end!")
	if number >= 5000 {
		fmt.Println("Too High...")
		return 5000
	}
	fmt.Println("Too low...")
	return number

}

func main() {
	fmt.Println(getApprovedValue(100))
}
