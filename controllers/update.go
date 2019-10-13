package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lkhadi/MVC_golang-practice/models"
	"github.com/lkhadi/MVC_golang-practice/views"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://testinglocal")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.WriteHeader(http.StatusOK)
		if r.Method == http.MethodPut {
			data := views.UpdateArtikel{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := models.Update(data); err != nil {
				w.Write([]byte("Cannot update data"))
				return
			}
			data2 := views.Response{
				Code: 200,
				Body: "Data Updated!",
			}
			json.NewEncoder(w).Encode(data2)
		} else {
			data := views.Response{
				Code: 500,
				Body: "Something Wrong",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}
