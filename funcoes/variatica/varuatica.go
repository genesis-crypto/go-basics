package main

import "fmt"

func med(numbers ...float64) float64 {
	sum := 0.0
	for _, value := range numbers {
		sum += value
	}

	return sum / float64(len(numbers))
}

func main() {
	fmt.Println(med(10.0, 20.0, 30.0))
}
