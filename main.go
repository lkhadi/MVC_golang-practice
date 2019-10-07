package main

import (
	"net/http"

	"testapi/controllers"
)

func main() {
	mux := controllers.Route()
	http.ListenAndServe("localhost:80", mux)
}
