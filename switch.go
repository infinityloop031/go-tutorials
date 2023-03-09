package main

import "fmt"

func main() {
	var city string
	fmt.Scan(&city)

	switch city{

	case "New York":
		fmt.Println("Hello NewYork")
	
	case "Pakistan":
		fmt.Println("Hello Pakistan")
	
	case "Multan":
		fmt.Println("Hello Multan")
	
	case "Mexico":
		fmt.Println("Hello Mexico")
	
	default:
		fmt.Println("No valid city selected")

	}
}
