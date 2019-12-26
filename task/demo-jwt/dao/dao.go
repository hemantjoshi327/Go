package dao

import(
	"database/sql"
)

func validateuser(){
	
	result, err := db.Query("SELECT id, name, city from company")
}