package main

import (
	"Account/AccountServ/database"
	"log"
)

func main() {
	db := database.MysqlDB
	// fmt.Println(db)
	log.Println(db)
}
