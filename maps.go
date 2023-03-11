package main

import (
	"fmt"
	"reflect"
)

func main() {

	var userData = make(map[string]string)
	userData["firstname"]="bilal"
	userData["lastnane"]="shabbir"

	fmt.Println(userData)

	fmt.Println(reflect.TypeOf(userData))
}
