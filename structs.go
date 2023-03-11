package main

import (
	"fmt"
	"reflect"
)

func main() {


	type UserData struct {
			firstName string
			lastName string
			phoneNo int
	}

	var user= UserData{
		firstName: "bilal",
		lastName: "shabbir",
		phoneNo:787844545 ,
	}

	fmt.Println(user)
	fmt.Println(reflect.TypeOf(user))
}
