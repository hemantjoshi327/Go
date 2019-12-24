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
//for get all the Employees

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	db := con.Connect()
	defer db.Close()
  
	w.Header().Set("Content-Type", "application/json")
	var employee []models.Employee
	result, err := db.Query("select employee.id, employee.name,employee.city,company.name AS company_name from employee join company where employee.company_id = company.id")
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
  
  //get employee by id
  
  func GetEmployee(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
	  w.Header().Set("Content-Type", "application/json")
	  params := mux.Vars(r)
	  result, err := db.Query("select employee.id, employee.name,employee.city,company.name AS company_name from employee join company where employee.company_id = company.id AND employee.id = ? ", params["id"])
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
  
  //create company
  
  func CreateEmployee(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
	  w.Header().Set("Content-Type", "application/json")
	  stmt, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
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
	  fmt.Fprintf(w, "New Employee was created")
	}
  
  //update company by id
  
  func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
	  w.Header().Set("Content-Type", "application/json")
	  params := mux.Vars(r)
	  stmt, err := db.Prepare("UPDATE employee SET name = ?, city = ? WHERE id = ?")
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
	  fmt.Fprintf(w, "Employee with ID = %s was updated", params["id"])
	}
  
  
  //delete employee by id 
  
  func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
  
	  db := con.Connect()
	  defer db.Close()
  
	  w.Header().Set("Content-Type", "application/json")
	  params := mux.Vars(r)
	  stmt, err := db.Prepare("DELETE FROM employee WHERE id = ?")
	  if err != nil {
		panic(err.Error())
	  }
	  _, err = stmt.Exec(params["id"])
	  if err != nil {
		panic(err.Error())
	  }
	  fmt.Fprintf(w, "Employee with ID = %s was deleted", params["id"])
	}