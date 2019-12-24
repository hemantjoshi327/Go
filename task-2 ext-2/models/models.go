package models

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
