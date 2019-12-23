package routes

import(
	s "../services"
	"github.com/gorilla/mux"
	"net/http"
	"log"


)

func Start(){

	log.Println("Server started on: http://localhost:8000")
	
	r := mux.NewRouter()
	api := r.PathPrefix("/company").Subrouter()
	api.HandleFunc("", s.GetCompanies).Methods("GET")
	api.HandleFunc("/{id}", s.GetCompany).Methods("GET")
	api.HandleFunc("", s.CreateCompany).Methods("POST")
	api.HandleFunc("/{id}", s.UpdateCompany).Methods("PUT")
	api.HandleFunc("/{id}", s.DeleteCompany).Methods("DELETE")
	http.ListenAndServe(":8000", r)


}