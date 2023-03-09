package main

import (
	"fmt"
	"strings"
)

func main(){
	bookings:= []string{}
	bookings=append(bookings,"hello sir")
	fmt.Println(bookings[0])
	firstnames :=[]string {}

	for _,booking := range bookings{
		var names =strings.Fields(booking)
		firstnames=append(firstnames,names[0])
	}
	fmt.Println(firstnames[0])
}


