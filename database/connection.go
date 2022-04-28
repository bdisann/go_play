package database

import "fmt"

var connection string

func init() {
	connection = "PostgreeSQL"
	fmt.Println("Connection Init")
}

func GetDB() string {
	return connection
}
