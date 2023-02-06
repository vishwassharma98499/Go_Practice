package main

import (
	"fmt"
)

var pl = fmt.Println

func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

// closure are function that dont have to be associated with an identifier(anomonous)
func main() {
	intSum := func(x, y int) int { return x + y }

	pl("5+ 4 =", intSum(5, 4))

	sampl := 1
	changeVar := func() { sampl += 1 }

	changeVar()
	pl("sampl1 =", sampl)

	useFunc(sumVals, 5, 8)
}
