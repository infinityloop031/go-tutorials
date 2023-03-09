package main

import "fmt"

func main() {
	fmt.Println("Main function")
	greetUser()
	para("bilal", 6)
	y:= mult(6)
	fmt.Println(y)
}

func greetUser() {
	fmt.Println("wlc to conference!")
}

func para(x string, b int) {
	fmt.Printf("hello %v, %v\n", x, b)
}

func mult(a int) int {
	return a*2
}
