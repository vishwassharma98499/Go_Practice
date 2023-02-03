package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var arr [8]int
	arr[0] = 1
	pl(arr)
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl(arr2)
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	arr3 := [2][2]int{
		{1, 2}, {10, 20},
	}
	pl(arr3)
	var i int
	pl(i)

	astr1 := "abcde"
	rArr := []rune(astr1)
	for _, v := range rArr {
		fmt.Printf("val %d", v)
	}
}
