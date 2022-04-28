package main

import (
	"fmt"
	"go_play/helper"
)

type Person struct {
	name string
}

func changeName(p *Person) {
	(*p).name = "santoso"
}

func main() {
	person := Person{
		name: "budi",
	}

	changeName(&person)

	fmt.Println(person)
	helper.SayHello(person.name)
}
