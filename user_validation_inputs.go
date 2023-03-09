package main

import (
	"fmt"
	"strings"
)

func main(){
	var firstname string
	fmt.Scan(&firstname)
	fmt.Println(firstname)
	isvalidName:= len(firstname)<5
	fmt.Println(isvalidName)
	email :="anything@gmail.com"
	isvalidEmail:=strings.Contains(email,"@")
	fmt.Println(isvalidEmail)
}