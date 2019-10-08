package controllers

import (
	"encoding/json"
	"net/http"
	"testapi/views"
)

func Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Access-Control-Allow-Origin", "http://testinglocal")
			w.WriteHeader(http.StatusOK)
			arr := [][]string{{"1", "Lalu"}, {"2", "Kismara"}, {"3", "Hadi"}}
			data := views.Response{
				Code: 200,
				Body: arr,
			}
			json.NewEncoder(w).Encode(data)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			data := views.Response{
				Code: 500,
				Body: "Something Wrong",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}
