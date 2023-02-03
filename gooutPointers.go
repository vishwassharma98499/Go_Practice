package main

import (
	"fmt"
)

var pl = fmt.Println

func chnageNum(n *float64) {

	*n = float64(*n + 1)

}

func doubleArray(arr *[4]int) {

	for x := 0; x < 4; x++ {
		arr[x] = arr[x] * 2
	}

}

func sumSlice(nums ...float64) float64 {

	var nsum float64 = 0
	var nSize float64 = float64(len(nums))
	pl(nSize)
	for _, val := range nums {
		nsum += val
	}

	return nsum
}
func main() {
	var num float64 = 2
	var pnum *float64 = &num

	pl(num)
	pl(pnum)

	chnageNum(&num)

	pl(num)
	pl(pnum)

	parr := [4]int{1, 2, 3, 4}

	pl(parr)
	doubleArray(&parr)
	pl(parr)

	pl(sumSlice(1, 2, 3, 4, 5))

	iSlice := []float64{1, 2, 3, 4}

	pl(sumSlice(iSlice...))
}
