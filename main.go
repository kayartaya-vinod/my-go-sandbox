package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	log.Println("Opening the connectoion")
	db, err := sql.Open("mysql", "root:Welcome#123@tcp(127.0.0.1)/northwind")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	data, err := db.Query("select * from SHIPPERS")
	if err != nil {
		panic(err)
	}

	for data.Next() {
		var id int
		var name, phone string
		data.Scan(&id, &name, &phone)
		log.Println(id, name, phone)
	}
}
