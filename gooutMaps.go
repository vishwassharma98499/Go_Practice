package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	pl("maps are collection that has key value pair")

	//var myMap map [keyType]valueType
	//var heroes map[string]string

	var heroes = make(map[string]string)
	villians := make(map[string]string)

	heroes["Batman"] = "bruce wayne"
	heroes["superman"] = "clark kent"
	heroes["flash"] = "barry allen"
	villians["lex luther"] = "lex luther"
	villians["superman"] = "clark kent"
	superPets := map[int]string{1: "Krypto", 2: "bat hound"}

	fmt.Printf("Batman is %v\n", heroes["Batman"])

	pl("chip :", superPets[3])

	_, ok := superPets[3]
	pl("is there a 3rd pet:", ok)

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	delete(heroes, "flash")
	pl(heroes)
}
