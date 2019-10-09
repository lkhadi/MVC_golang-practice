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
		if r.Method == http.MethodPost {
			data := views.Artikel{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := models.Artikel(data.Artikel, data.Judul, data.Author, data.Tanggal); err != nil {
				w.Write([]byte("Cannot insert data"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			data2 := views.Response{
				Code: 201,
				Body: "Success",
			}
			json.NewEncoder(w).Encode(data2)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			data2 := views.Response{
				Code: 500,
				Body: "Something Wrong",
			}
			json.NewEncoder(w).Encode(data2)
		}
	}
}
