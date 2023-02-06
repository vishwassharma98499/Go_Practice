package stuff

import "strconv"

var Name string = "Vishwas"

func IntArraytoStrArray(intArr []int) []string {
	var strArr []string

	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr

}

// capital N in var name and capital I in function name signifies it can be accesed outside tis file
