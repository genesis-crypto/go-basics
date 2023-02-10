package main

import "fmt"

func main() {
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 2.1, 3.1, 4.1
	fmt.Println(grades)

	var total float64
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	avrage := total / float64(len(grades))
	fmt.Printf("Avrage: %.2f", avrage)
}
