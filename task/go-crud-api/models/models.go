package models

import(
	"github.com/dgrijalva/jwt-go"
)

type Company struct {
	ID string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"` 
}

type Employee struct{
	ID int `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Company string `json:"company"`
}  


type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}