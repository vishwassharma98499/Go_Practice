package main

import (
	"fmt"
)

var pl = fmt.Println

type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {

	ml1 := ML(Tsp(3) * 4.92)

	fmt.Printf("3 tsps = %.2f ML \n", ml1)

	mL2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 tsps = %.2f ML \n", mL2)

	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp >4 tsp =", Tsp(2) > Tsp(4))

	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 Tbs = %.2f mL\n", TBToML(3))

	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
}
