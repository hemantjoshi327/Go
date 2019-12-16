package main

import "fmt"

func sub(x int, y int) int {
	return x - y
}

func add(x int, y int) int {
	return x + y
}

func div(x int, y int) int {
	return x / y
}

func rem(x int, y int) int {
	return x % y
}

func mul(x int, y int) int {
	return x * y
}


func main() {
	fmt.Println(add(2,3))
	fmt.Println(sub(3,3))
	fmt.Println(div(10,5))
	fmt.Println(rem(10,3))
	fmt.Println(mul(10,3))
}
