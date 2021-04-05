package main

import (
    "net/http"
    "github.com/TheOnly-Co/to_do/controller"
    "github.com/TheOnly-Co/to_do/model"
    "log"
    "fmt"
)



func main() {
    mux := controller.Register()
    
    db := model.Connect()
    defer db.Close()
    fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
