package main

import (
	"fmt"
	"strings"
)

type People struct {
	name     string
	forename string
}

func (p People) getFullName() string {
	return p.name + " " + p.forename
}

func (p *People) setFullName(completeName string) {
	names := strings.Split(completeName, " ")
	p.name = names[0]
	p.forename = names[1]
}

func main() {
	first_person := People{name: "Lo", forename: "Po"}
	fmt.Println(first_person.getFullName())
	first_person.setFullName("Othe Sla")
	fmt.Println(first_person.getFullName())
}
