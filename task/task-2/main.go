package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
  )

  type Company struct {
	ID string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
  }

	var db *sql.DB
	var err error

//for get all the companies

func getCompanies(w http.ResponseWriter, r *http.Request) {
	
  w.Header().Set("Content-Type", "application/json")
  var company []Company
  result, err := db.Query("SELECT id, name, city from company")
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var cmp Company
    err := result.Scan(&cmp.ID, &cmp.Name, &cmp.City)
    if err != nil {
      panic(err.Error())
    }
    company = append(company, cmp)
  }
  json.NewEncoder(w).Encode(company)
}

//get company by id

func getCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, name, city FROM company WHERE id = ?", params["id"])
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	var cmp Company
	for result.Next() {
	  err := result.Scan(&cmp.ID, &cmp.Name, &cmp.City)
	  if err != nil {
		panic(err.Error())
	  }
	}
	json.NewEncoder(w).Encode(cmp)
  }

//create company

func createCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO company(name, city) VALUES(?,?)")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	city := keyVal["city"]

	_, err = stmt.Exec(name, city)
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "New company was created")
  }

//update company by id

func updateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE company SET name = ?, city = ? WHERE id = ?")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["name"]
	newCity := keyVal["city"]
	_, err = stmt.Exec(newName, newCity, params["id"])
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Company with ID = %s was updated", params["id"])
  }


//delete company by id 

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM company WHERE id = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Company with ID = %s was deleted", params["id"])
  }



func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/goblog")
	  if err != nil {
		panic(err.Error())
	  }

	  
	  defer db.Close()
	 
	  router := mux.NewRouter()
	  router.HandleFunc("/company", getCompanies).Methods("GET")
	  router.HandleFunc("/company/{id}", getCompany).Methods("GET")
	  router.HandleFunc("/company", createCompany).Methods("POST")
	  router.HandleFunc("/company/{id}", updateCompany).Methods("PUT")
	  router.HandleFunc("/company/{id}", deleteCompany).Methods("DELETE")
	  log.Println("Server started on: http://localhost:8000")
	  http.ListenAndServe(":8000", router)
	}