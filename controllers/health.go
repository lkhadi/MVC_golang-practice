package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lkhadi/MVC_golang-practice/views"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://testinglocal")
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			data := views.Response{
				Code: 200,
				Body: "Running",
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
