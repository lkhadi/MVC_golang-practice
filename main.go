package main

import (
	"log"
	"net/http"
	"testapi/controllers"
	"testapi/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controllers.Route()
	db := models.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
