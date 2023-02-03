package main

import (
	"fmt"
)

var pl = fmt.Println

func sayHello() {
	pl("hello")
}

// radical function with unknown parameter
func getSum2(nums ...int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sum(x int, y int) int {

	return x + y
}

func getQuotiend(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("yoi cant devide by 0")
	} else {
		return x / y, nil
	}

}
func main() {
	// func functionName (PARAMETER) returnType {Body}
	sayHello()
	pl(sum(1, 2))

	ans, err := getQuotiend(2, 0)
	if err != nil {
		pl(err)
	} else {
		fmt.Printf("%d / %d = ", 0, 2, ans)
	}

	ans, err = getQuotiend(2, 2)
	if err != nil {
		pl(err)
	} else {
		fmt.Printf("%d / %d = %f", 2, 2, ans)
	}
	ans1, err1 := getQuotiend(2, 2)
	if err1 != nil {
		pl(err1)
	} else {
		fmt.Printf("%d / %d = %f", 2, 2, ans1)
	}
	pl("/n")
	pl(getSum2(1, 2, 3))
}
