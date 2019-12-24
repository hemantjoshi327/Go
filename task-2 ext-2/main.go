package main

import (
	_"github.com/go-sql-driver/mysql"
	"./routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
  )


func main() {	
	log.Println("Server started on: http://localhost:8000")
	r := mux.NewRouter()
	routes.Company(r)
	routes.Employee(r)
	routes.Credentials(r)
	log.Fatal(http.ListenAndServe(":8000", r)) 
	}

	