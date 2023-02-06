package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)

	intArr := []int{2, 3, 4, 5, 6, 11}
	strArr := stuff.IntArraytoStrArray(intArr)

	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

}
