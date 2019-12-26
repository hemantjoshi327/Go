package controller

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	dao "../dao"
  )
  
  // Create the JWT key used to create the signature
  var jwtKey = []byte("my_secret_key")
  
  var users = map[string]string{
	  "Hemant": "JJJJJ",
	  "Harry": "PPPPP",
  }
  
  // Create a struct to read the username and password from the request body
  type Credentials struct {
		Username string `json:"username"`  
		Password string `json:"password"`
	  
  }
  
  // Create a struct that will be encoded to a JWT.
  type Claims struct {
	  Username string `json:"username"`
	  jwt.StandardClaims
  }
  
  // Create the Login handler
  func Signin(w http.ResponseWriter, r *http.Request) {
	  var creds Credentials
	  err := json.NewDecoder(r.Body).Decode(&creds)
	  if err != nil {
		  w.WriteHeader(http.StatusBadRequest)
		  return
	  }
	  	user 
	  	validCreds = dao.LoginValid(&creds)
	 // expectedPassword, ok := users[creds.Username]
  
	  if !ok || expectedPassword != creds.Password {
		  w.WriteHeader(http.StatusUnauthorized)
		  return
	  }
  
	  // 5 minutes
	  expirationTime := time.Now().Add(5 * time.Minute)
	  // JWT claims
	  claims := &Claims{
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