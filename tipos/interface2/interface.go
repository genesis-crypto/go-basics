package main

import "fmt"

type Sporty interface {
	turnOn()
}

type Ferrari struct {
	model        string
	onTurbo      bool
	currentSpeed int
}

func (f *Ferrari) turnOn() {
	f.onTurbo = true
}

func main() {
	la_ferrari := Ferrari{"La Ferrari", false, 0}
	la_ferrari.turnOn()

	var romeo Sporty = &Ferrari{"Romeo", false, 0}
	romeo.turnOn()

	fmt.Println(la_ferrari.onTurbo, romeo)
}
