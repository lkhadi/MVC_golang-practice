package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lkhadi/MVC_golang-practice/models"
	"github.com/lkhadi/MVC_golang-practice/views"
)

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://testinglocal")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE")
		w.WriteHeader(http.StatusOK)
		if r.Method == http.MethodDelete {
			idartikel := r.URL.Path[8:]
			err := models.Delete(idartikel)
			if err != nil {
				w.Write([]byte("Cannot read data"))
				return
			}
			data := views.Response{
				Code: 200,
				Body: "Data deleted",
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
