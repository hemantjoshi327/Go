// package controller

// import(
// 	con "../config"		
// 	s "../services"
// 	"net/http")

// func GetCompanies(w http.ResponseWriter, r *http.Request) {
// 	db := con.Connect()
// 	defer db.Close()
  
// 	w.Header().Set("Content-Type", "application/json")
// 	company := s.Getall()

// 	json.NewEncoder(w).Encode(company)
//   }