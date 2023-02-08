package main

import (
	"fmt"
	"time"
)

func typeOf(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Real"
	case func():
		return "Function"
	case string:
		return "String"
	case bool:
		return "Boolean"
	default:
		return "not know"
	}
}

func main() {
	fmt.Println(typeOf(2))
	fmt.Println(typeOf(true))
	fmt.Println(typeOf(func() {}))
	fmt.Println(typeOf("a"))
	fmt.Println(typeOf(time.Now()))
}
