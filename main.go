package main
import (
    "os"
	"net/http"
)


func main() {
		http.HandleFunc("/",handler)
		http.ListenAndServe(os.Agrs[1], nil)
}

func handler(w http.ResponseWriter, req *http.Request){
	query := req.URL.Query()
	param := query.Get("")


	if req.Method == MethodPut{
		

	}
	if req.Method == MethodGet{
		param := query.Get("timeout")

	}



}