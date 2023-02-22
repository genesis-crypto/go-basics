package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	shoes := product{
		name:     "Nike",
		price:    1000,
		discount: 0.5,
	}
	fmt.Printf("Shoes Price > %.2f => DISCOUNT %.2f\n", shoes.price, shoes.priceWithDiscount())

	computer := product{"Dell", 10000, 0.1}
	fmt.Printf("Computer Price > %.2f => DISCOUNT %.2f\n", computer.price, computer.priceWithDiscount())

}
