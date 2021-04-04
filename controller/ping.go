package controller


import (
	"net/http"
    "encoding/json"
    "github.com/TheOnly-Co/to_do/views"
)
func ping() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
        if r.Method == http.MethodGet {
        	data := structs.Response{
        		Code: http.StatusOK,
        		Body: "pong",
        	}
        	json.NewEncoder(w).Encode(data)
        	fmt.Println("recieved")
        }
	}
}