package main

import(
	"fmt"
)

func main(){

	// 1 doller = 70.17
	var doller, rupees float32 
	doller = 70.17
	
	fmt.Println("Enter a amount in rupees to convert into US Doller")
	fmt.Scan(&rupees)

	fmt.Println("You entered rupeess",rupees)
	doller = rupees / doller 
	fmt.Println("Your money in doller = "+"$",doller)

}