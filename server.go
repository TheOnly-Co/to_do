package main

import (
    "net/http"
    "github.com/TheOnly-Co/to_do/views"
    "encoding/json"
)



func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/ping",func(w http.ResponseWriter, r *http.Request){
        if r.Method == http.MethodGet {
        	data := views.Response{
        		Code: http.StatusOK,
        		Body: "pong",
        	}
        	json.NewEncoder(w).Encode(data)
        	
        }
	})
   
	http.ListenAndServe("localhost:3000", mux)
}
