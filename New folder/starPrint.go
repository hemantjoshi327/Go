package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to go tour")
	no := 5
	// alpha(no)
	// fmt.Println()
	// revAlpha(no)
	// fmt.Println()
	// star(no)
	// fmt.Println()
	// revStar(no)
	starRev(no)

}

// func alpha(n int) {

// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= i; j++ {

// 			fmt.Print(j)
// 		}
// 		fmt.Println()
// 	}
// }

// func revAlpha(n int) {

// 	for i := n; i >= 1; i-- {
// 		for j := 1; j <= i; j++ {

// 			fmt.Print(j)
// 		}
// 		fmt.Println()
// 	}
// }

// func star(n int) {

// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= i; j++ {

// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func revStar(n int) {

// 	for i := n; i >= 1; i-- {
// 		for j := 1; j <= i; j++ {

// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }


func starRev(n int){
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {

			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i := n-1; i >= 1; i-- {
		for j := 1; j <= i; j++ {

			fmt.Print("* ")
		}
		fmt.Println()
	}
}
