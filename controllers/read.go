package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lkhadi/MVC_golang-practice/models"
	"github.com/lkhadi/MVC_golang-practice/views"
)

func Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Access-Control-Allow-Origin", "http://testinglocal")
			w.WriteHeader(http.StatusOK)
			arr, err := models.ReadAll()
			if err != nil {
				w.Write([]byte("Cannot read data"))
			}
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

func ReadOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://testinglocal")
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			idartikel := r.URL.Query().Get("id")
			arr, err := models.ReadOne(idartikel)
			if err != nil {
				w.Write([]byte("Cannot read data"))
			}
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
