package main

import "fmt"

type Car struct {
	name         string
	currentSpeed float64
}

type Ferrari struct {
	Car
	turboActive bool
}

func main() {
	la_ferrari := Ferrari{}
	la_ferrari.name = "La Ferrari"
	la_ferrari.currentSpeed = 180
	la_ferrari.turboActive = true
	fmt.Println(la_ferrari)
}
