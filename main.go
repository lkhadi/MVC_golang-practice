package main

import (
	"log"
	"net/http"
	"github.com/lkhadi/MVC_golang-practice/controllers"
	"github.com/lkhadi/MVC_golang-practice/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controllers.Route()
	db := models.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
