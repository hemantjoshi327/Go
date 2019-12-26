package dao

import(
	"../models"
	con "../config"
)

func LoginValid() {
	db := con.Connect()
	defer db.Close()

	var usr []models.User
	result, err := db.Query("SELECT name, password from user where name = ? AND password = ?",)
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  err := result.Scan(&usr.Uname, &usr.Password)
	  if err != nil {
		panic(err.Error())
	  }
	}
  }