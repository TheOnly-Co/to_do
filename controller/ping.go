package controller

import(
	"net/http"
	"github.com/TheOnly-Co/to_do/views"
    "encoding/json"
)
func ping() http.HandleFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		
		}
	}
}