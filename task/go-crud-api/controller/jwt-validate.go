package controller

import (
	// "fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"time"
	"../models"
	con "../config"
)

var jwtKey = []byte("my_secret_key")


func Signin(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//database
	db := con.Connect()
	defer db.Close()

	result, err := db.Query("SELECT name, passwordl FROM users WHERE name = ?", creds.Username)
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	var vcreds models.Credentials
	for result.Next() {
	  err := result.Scan(&vcreds.Username, &vcreds.Password)
	  if err != nil {
		panic(err.Error())
	  }
	}
	// fmt.Println(vcreds.Password)

	expectedPassword :=  vcreds.Password
	// fmt.Println(expectedPassword)


	if expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}