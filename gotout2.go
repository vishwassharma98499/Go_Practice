package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

// var start with capital mean exported or available outside the pacage
func main() {

	//var vName string = "Vishwas"
	//var v1, v2 = 1.2, 3.4
	//var v3 = "Hello"
	//v1 = 4.5
	// int , float63,bool,string,rund
	//Default type 0,0.0,false,""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(2.5))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	cv1 := 1.5
	cv2 := int(cv1) // typecast
	cv3 := "2"
	cv4, err := strconv.Atoi(cv3)
	cv5 := 50000
	cv6 := strconv.Itoa(cv5) //convert integer to string
	pl(cv4, err, reflect.TypeOf(cv4))
	pl(cv1)
	pl(cv2)
	pl(cv3)
	pl(cv6)

	if cv1 < 1.5 {
		pl("less than 1")
	} else if cv1 == 1.5 {
		pl("equal ")
	}

	fmt.Printf("%s %f %d %t %o \n", "string", 1.1, 1, true, 1) // define the integer types
}
