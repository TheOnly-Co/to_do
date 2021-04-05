package model 

import (
	"database/sql"
	"log"
	"fmt"

)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	return db

}