package main

import (
    "net/http"
    "github.com/TheOnly-Co/to_do/controller"
)



func main() {
    mux := controller.Register()
    
   
	http.ListenAndServe("localhost:3000", mux)
}
