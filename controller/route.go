package controller


import(

	"net/http"
	"github.com/TheOnly-Co/to_do/views"
    "encoding/json"
)
func Register() *http.ServeMux {
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
	return mux

}