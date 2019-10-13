package main

import (
	"log"
	"net/http"
	"github.com/lkhadi/MVC_golang-practice/controllers"
	"github.com/lkhadi/MVC_golang-practice/models"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	mux := controllers.Route()
	db := models.Connect()
	defer db.Close()
	fmt.Println("Running....")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
