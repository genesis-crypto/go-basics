package main

import (
	"fmt"
	"reflect"
	"strconv"
	c "strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// ASCII
	fmt.Println("ASCII: ", string(123))

	//int <> string
	fmt.Println("Int <> String: ", c.Itoa(123))

	// string <> int
	num, _ := c.Atoi("123")
	fmt.Println("String <> Int: ", num, reflect.TypeOf(num))

	b, err := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	} else {
		fmt.Println(err)
	}
}
