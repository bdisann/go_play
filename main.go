package main

import (
	// "fmt"
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type Persons []Person

func (self Persons) Len() int {
	return len(self)
}

func (self Persons) Less(i, j int) bool {
	return self[i].age < self[j].age
}

func (self Persons) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func main() {
	users := Persons{
		Person{name: "a", age: 5},
		Person{name: "b", age: 2},
		Person{name: "d", age: 3},
	}
	fmt.Println(users)
	sort.Sort(Persons(users))
	fmt.Println(users)
}
