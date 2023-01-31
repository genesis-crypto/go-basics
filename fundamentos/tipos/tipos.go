package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Int 8, 16, 32, 64
	fmt.Println(1, 2, -1000)
	fmt.Printf("\nTypeOf %d is %s", 32000, reflect.TypeOf(32000))

	// Uint 8, 16, 32, 64
	var a byte = 1
	fmt.Printf("\nByte is %s", reflect.TypeOf(a))

	i1 := math.MaxInt
	fmt.Println("O valor max de int é", i1)
	fmt.Println("O valor de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // represneta o valor na table unicode
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Float 8, 16, 32, 64
	var x float32 = 49.99
	fmt.Println("TypeOf x is", reflect.TypeOf(x))
	fmt.Println("Literal TypeOf is", reflect.TypeOf(42.22))

	// Bool
	var b bool = true
	fmt.Println("Typeof b is", reflect.TypeOf(b))
	fmt.Println(!b)

	// String
	var s string = "Hello, World"
	fmt.Println(s + "!")
	fmt.Println("String size is", len(s))

	// String multline
	var s2 string = `Hello
	World`
	fmt.Println("Size of s2 is", len(s2))

	// Char
	var c = 'a'
	fmt.Println("Type char is", reflect.TypeOf(c))
	fmt.Println(c)
}
