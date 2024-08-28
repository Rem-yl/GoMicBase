package main

import (
	"Store/StoreServ/database"
	"fmt"
)

func main() {
	db := database.MysqlDB

	fmt.Println(db)
}
