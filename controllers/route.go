package controllers

import "net/http"

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health())
	mux.HandleFunc("/read", Read())
	mux.HandleFunc("/create", Create())
	return mux
}
