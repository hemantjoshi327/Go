package controller

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
  )
  
  // Create the JWT key used to create the signature
  var jwtKey = []byte("my_secret_key")
  
  var users = map[string]string{
	  "user1": "password1",
	  "user2": "password2",
  }
  
  // Create a struct to read the username and password from the request body
  type Credentials struct {
	  Password string `json:"password"`
	  Username string `json:"username"`
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
  
	  expectedPassword, ok := users[creds.Username]
  
	  if !ok || expectedPassword != creds.Password {
		  w.WriteHeader(http.StatusUnauthorized)
		  return
	  }
  
	  // here, we have kept it as 5 minutes
	  expirationTime := time.Now().Add(5 * time.Minute)
	  // Create the JWT claims, which includes the username and expiry time
	  claims := &Claims{
		  Username: creds.Username,
		  StandardClaims: jwt.StandardClaims{
			  ExpiresAt: expirationTime.Unix(),
		  },
	  }
  
	  // Declare the token with the algorithm used for signing, and the claims
	  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	  // Create the JWT string
	  tokenString, err := token.SignedString(jwtKey)
	  if err != nil {
		  // If there is an error in creating the JWT return an internal server error
		  w.WriteHeader(http.StatusInternalServerError)
		  return
	  }

	  http.SetCookie(w, &http.Cookie{
		  Name:    "token",
		  Value:   tokenString,
		  Expires: expirationTime,
	  })
  }