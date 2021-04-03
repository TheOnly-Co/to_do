package main

import (
    "net/http"
    "fmt"
)

{
    "key"
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
        if r.Method == http.MethodGet {
        	struct
        }
    })
   
	http.ListenAndServe("localhost:3000", nil)
}
