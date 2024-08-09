package main

import (
	"Account/Database"
	"fmt"
)

func main() {
	db := Database.MysqlDB
	fmt.Println(db)
}
