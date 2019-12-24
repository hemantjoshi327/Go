package routes

import(
	cont "../controller"
	"github.com/gorilla/mux"


)


var Company = func(router *mux.Router){	
	cmp := router.PathPrefix("/company").Subrouter()
	cmp.HandleFunc("", cont.GetCompanies).Methods("GET")
	cmp.HandleFunc("/{id}", cont.GetCompany).Methods("GET")
	cmp.HandleFunc("", cont.CreateCompany).Methods("POST")
	cmp.HandleFunc("/{id}", cont.UpdateCompany).Methods("PUT")
	cmp.HandleFunc("/{id}", cont.DeleteCompany).Methods("DELETE")
	//employee CRUD by company id and emp id
	cmp.HandleFunc("/{id}/employee", cont.GetCmpEmployees).Methods("GET")
	cmp.HandleFunc("/{cid}/employee/{eid}", cont.GetCmpEmployee).Methods("GET")
}

var Employee = func(router *mux.Router){
	emp := router.PathPrefix("/employee").Subrouter()
	emp.HandleFunc("", cont.GetEmployees).Methods("GET")
	emp.HandleFunc("/{id}", cont.GetEmployee).Methods("GET")
	emp.HandleFunc("", cont.CreateEmployee).Methods("POST")
	emp.HandleFunc("/{id}", cont.UpdateEmployee).Methods("PUT")
	emp.HandleFunc("/{id}", cont.DeleteEmployee).Methods("DELETE")
}

var Credentials = func(router *mux.Router){
	router.HandleFunc("/login", cont.Signin)
}


