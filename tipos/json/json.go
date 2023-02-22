package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	product_one := Product{1, "notebook", 12000, []string{"promo"}}
	procut_json, _ := json.Marshal(product_one)

	fmt.Println(string(procut_json))

	var product_two Product
	json_string := `{"id":1,"name":"notebook dell","price":2000,"tags":["promo"]}`
	json.Unmarshal([]byte(json_string), &product_two)
	fmt.Println(json_string)
}
