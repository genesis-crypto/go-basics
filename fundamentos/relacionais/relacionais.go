package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">= ", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type People struct {
		Name string
	}

	first_people := People{"Ke"}
	second_people := People{"Ke"}

	fmt.Println("First == Second? ", first_people == second_people)
}
