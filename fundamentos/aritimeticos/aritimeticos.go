package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("Sum => ", a+b)
	fmt.Println("Sub => ", a-b)
	fmt.Println("Div => ", a/b)
	fmt.Println("Mult => ", a*b)
	fmt.Println("Mod => ", a%b)

	//bitwise
	fmt.Println("Bit => ", a&b)
	fmt.Println("Or => ", a|b)
	fmt.Println("Xor => ", a^b)

	// math
	c := 3.0
	d := 2.0

	fmt.Println("Pow =>", math.Pow(c, d))
	fmt.Printf("Max => %v <> Min => %v", math.Max(c, d), math.Min(c, d))
}
