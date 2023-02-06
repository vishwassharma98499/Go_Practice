package main

import (
	"fmt"
)

var pl = fmt.Println

type customer struct {
	name    string
	address string
	balance float64
}

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("contact at %s is %s %s\n",
		b.name, b.contact, b.contact.phone)

}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
}

func newCustAdd(c *customer, address string) {
	c.address = address

}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

func main() {

	var tS customer
	tS.name = "Tom smith"
	tS.address = "alter postweg"
	tS.balance = 234.56

	getCustInfo(tS)

	newCustAdd(&tS, "123 address")
	pl("addres :", tS.address)

	sS := customer{"sally smith", "123 main", 0.0}

	pl("name :", sS.name)

	rect1 := rectangle{10.0, 15.0}
	pl("rect area :", rect1.Area())

	con1 := contact{"james", "wang", "555-1212"}

	bus1 := business{"abc", "234 north street", con1}

	bus1.info()
}
