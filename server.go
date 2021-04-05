package main

import (
    "net/http"
    "github.com/TheOnly-Co/to_do/controller"
    "github.com/TheOnly-Co/to_do/model"
    "github.com/go-sql-driver/mysql"
)



func main() {
    mux := controller.Register()
    
    db := model.Connect()
    defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
