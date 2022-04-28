package main

import (
	"fmt"
	"go_play/database"
)

// type Person struct {
// 	name string
// }

// func changeName(p *Person) {
// 	(*p).name = "santoso"
// }

func main() {
	dbName := database.GetDB()

	fmt.Println(dbName)
}
