package main

import "fmt"

func main() {
	var approved map[int]string = make(map[int]string)

	approved[12321312] = "A"
	approved[12322522] = "V"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%v - %v\n", cpf, name)
	}

	fmt.Printf("> %v\n", approved[12321312])
	delete(approved, 12322522)
	fmt.Printf("Deleted: %v", approved[12322522])

}
