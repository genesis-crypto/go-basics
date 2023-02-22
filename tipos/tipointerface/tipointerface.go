package main

import "fmt"

type Course struct {
	name string
}

func main() {
	var some interface{}
	fmt.Println(some)

	some = 3
	fmt.Println(some)
	some = "hello"

	type dynamicType interface{}

	var aa dynamicType = "yeah"

	fmt.Println(aa)
	fmt.Println(some)

}
