package main

import "fmt"

type Show interface {
	toString() string
}

type People struct {
	name     string
	forename string
}

type Product struct {
	name  string
	price float64
}

func (p People) toString() string {
	return p.name + " " + p.forename
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func show(x Show) {
	fmt.Println(">> ", x.toString())
}

func main() {
	first_person := People{
		name:     "Ja",
		forename: "La",
	}

	show(first_person)
}
