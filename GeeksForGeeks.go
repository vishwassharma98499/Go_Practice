package main

import (
	"fmt"
)

var pl = fmt.Println

// var start with capital mean exported or available outside the pacage
func main() {

}

/*\\
func findLongestConseqSubseq(arr, N) {
	arr := sorted(arr)
	nNumSize := 1
	nPrevNumSize := 1
	nPrev := 0
	fmt.Println(arr)
	for n := 1; n < N; n++ {
		if arr[n-1] == arr[n] {
			continue
		}
		fmt.Println(arr[n])
		if (arr[n-1] + 1) == arr[n] {
			nNumSize += 1
		}
		else {
			nNumSize = 1
			if nPrevNumSize < nNumSize {
				nPrevNumSize = nNumSize
			}
		}

		fmt.Print.ln(arr[n-1], arr[n], nNumSize)
	}
	return nPrevNumSize

}
*/
