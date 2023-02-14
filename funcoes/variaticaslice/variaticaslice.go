package main

import "fmt"

func showApproved(approved ...string) {
	for position, person := range approved {
		fmt.Printf("%d - %s\n", position+1, person)
	}
}

func main() {
	approvedPeople := []string{"ka", "aa", "ba"}

	showApproved(approvedPeople...)
}
