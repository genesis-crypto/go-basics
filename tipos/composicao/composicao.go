package main

import "fmt"

type Sporty interface {
	onTurbo()
}

type Luxous interface {
	doParallelPark()
}

type SportAndLuxous interface {
	Sporty
	Luxous
}

type AMG struct{}

func (c AMG) onTurbo() {
	fmt.Println("Tubo!")
}

func (c AMG) doParallelPark() {
	fmt.Println("Parallel Park!")
}

func main() {
	var amg SportAndLuxous = AMG{}

	amg.onTurbo()
	amg.doParallelPark()
}
