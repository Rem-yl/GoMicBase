package main

import (
	"Account/Database"
	"log"
)

func main() {
	db := Database.MysqlDB
	// fmt.Println(db)
	log.Println(db)
}
