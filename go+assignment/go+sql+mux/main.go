package main

import(
    "fmt"
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
)

//book struct

type employee struct{
	ID    int 	 `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"` 
}

type company struct{
	ID    int 	 `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"` 
}

// Database connection
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "goblog"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}


// get Companies
func getCompanies(w http.ResponseWriter, r *http.Request){

    db := dbConn()
    selDB, err := db.Query("SELECT * FROM company ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    cmp := company{}
    res := []company{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        cmp.ID = id
        cmp.Name = name
        cmp.City = city
        res = append(res, cmp)
    }
    defer db.Close()


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// get Company
func getCompany(w http.ResponseWriter, r *http.Request){

    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	fmt.Println(params)
	db := dbConn()
    selDB, err := db.Query("SELECT * FROM company WHERE id=?", params["id"])
    if err != nil {
        panic(err.Error())
    }
    cmp := company{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        cmp.ID = id
        cmp.Name = name
        cmp.City = city
    }
	
	json.NewEncoder(w).Encode(cmp)
    defer db.Close()

}


//create Company
func createCompany(w http.ResponseWriter, r *http.Request){

    

}

//update Company
// func updateBook(w http.ResponseWriter, r *http.Request){

// }

//delete Company
func deleteCompany(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    del := mux.Vars(r) //get delet
    fmt.Println(del)
    
    db := dbConn()
    delForm, err := db.Query("DELETE FROM company WHERE id=?", del["id"])
    if err != nil {
        panic(err.Error())
    }
    log.Println("DELETE")

    json.NewEncoder(w).Encode(delForm)
    defer db.Close()
}



func main(){
	//router
	r := mux.NewRouter()
	
	// Endpoints
    r.HandleFunc("/api/Company", getCompanies).Methods("GET")
    r.HandleFunc("/api/Company/{id}", getCompany).Methods("GET")
	// r.HandleFunc("/api/Company", createCompany).Methods("POST")
	// r.HandleFunc("/api/Company/{id}", updateCompany).Methods("PUT")
	r.HandleFunc("/api/Company/{id}", deleteCompany).Methods("Delete")
	
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}