package controllers

import "net/http"

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health())
	mux.HandleFunc("/read", Read())
	mux.HandleFunc("/create", Create())
	mux.HandleFunc("/read/artikel/", ReadOne())
	mux.HandleFunc("/delete/", Delete())
	return mux
}
