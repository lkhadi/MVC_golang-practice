package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lkhadi/MVC_golang-practice/models"
	"github.com/lkhadi/MVC_golang-practice/views"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://testinglocal")
		if r.Method == http.MethodGet {
			if err := models.Artikel(); err != nil {
				w.Write([]byte("Cannot insert data"))
				return
			}
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
