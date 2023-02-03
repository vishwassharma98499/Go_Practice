package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	sl1 := make([]string, 6)
	sl1[0] = "elem1"
	sl1[1] = "elem1"
	sl1[2] = "elem1"
	sl1[3] = "elem1"
	sl1[4] = "elem1"

	pl("size is", len(sl1))

	for i := 0; i < len(sl1); i++ {

		pl(sl1[i])
	}
	for _, x := range sl1 {

		pl(x)
	}

	sl2 := []int{1, 2, 3}
	pl(sl2)

	sl3 := sl2[:2]
	pl(sl3)
	sl2[0] = 10
	pl(sl2)
	pl(sl3)
	sl3 = append(sl3, 12)
	pl(sl3)
	pl(sl2)
}
