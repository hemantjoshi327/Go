package controller

import(
	"../models"
	con "../config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"

)
//for get all the companies

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	db := con.Connect()
	defer db.Close()
  
	w.Header().Set("Content-Type", "application/json")
	var company []models.Company
	result, err := db.Query("SELECT id, name, city from company")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var cmp models.Company
	  err := result.Scan(&cmp.ID, &cmp.Name, &cmp.City)
	  if err != nil {
		panic(err.Error())
	  }
	  company = append(company, cmp)
	}
	json.NewEncoder(w).Encode(company)
  }
  
  //get company by id
  
  func GetCompany(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
	  w.Header().Set("Content-Type", "application/json")
	  params := mux.Vars(r)
	  result, err := db.Query("SELECT id, name, city FROM company WHERE id = ?", params["id"])
	  if err != nil {
		panic(err.Error())
	  }
	  defer result.Close()
	  var cmp models.Company
	  for result.Next() {
		err := result.Scan(&cmp.ID, &cmp.Name, &cmp.City)
		if err != nil {
		  panic(err.Error())
		}
	  }
	  json.NewEncoder(w).Encode(cmp)
	}
  
  //create company
  
  func CreateCompany(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
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
  
  func UpdateCompany(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
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
  
  func DeleteCompany(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
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

	//get all employees from company by id

	func GetCmpEmployees(w http.ResponseWriter, r *http.Request) {
		db := con.Connect()
		defer db.Close()
	
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var employee []models.Employee
		result, err := db.Query("select employee.id, employee.name,employee.city,company.name AS company_name from employee join company where employee.company_id = company.id AND employee.company_id = ? ", params["id"])
		if err != nil {
			panic(err.Error())
		  }
		  defer result.Close()
		  for result.Next() {
			var emp models.Employee
			err := result.Scan(&emp.ID, &emp.Name, &emp.City, &emp.Company)
			if err != nil {
			  panic(err.Error())
			}
			employee = append(employee, emp)
		  }
		  json.NewEncoder(w).Encode(employee)
	  }

	//get employees from company by id

	func GetCmpEmployee(w http.ResponseWriter, r *http.Request) {
  
		db := con.Connect()
		defer db.Close()
	
		w.Header().Set("Content-Type", "application/json")
		cid := mux.Vars(r)["cid"]
		eid := mux.Vars(r)["eid"]
		result, err := db.Query("select employee.id, employee.name,employee.city,company.name AS company_name from employee join company where employee.company_id = company.id AND employee.company_id = ? AND employee.id = ?", cid, eid)
		if err != nil {
		  panic(err.Error())
		}
		defer result.Close()
		var emp models.Employee
		for result.Next() {
		  err := result.Scan(&emp.ID, &emp.Name, &emp.City, &emp.Company)
		  if err != nil {
			panic(err.Error())
		  }
		}
		json.NewEncoder(w).Encode(emp)
	  }  