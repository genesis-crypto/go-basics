package main

import "fmt"

type Item struct {
	id       int
	quantity int
	price    float64
}

type Order struct {
	user_id int
	items   []Item
}

func (order Order) totalValue() float64 {
	acc := 0.0
	for _, item := range order.items {
		acc += item.price * float64(item.quantity)
	}
	return acc
}

func main() {

	order := Order{user_id: 1, items: []Item{
		{id: 1, quantity: 2, price: 4.9},
	}}
	fmt.Println(order.totalValue())
}
