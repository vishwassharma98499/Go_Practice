package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	pl("enter number")
	reader := bufio.NewReader(os.Stdin) // stander input, buffer input output
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	s := strings.TrimSpace(name)
	ival, err := strconv.Atoi(s) // string to integer convert
	if err != nil {
		log.Fatal(err) // log the error
	}
	val1 := ival

	pl("enter Other Number")
	reader1 := bufio.NewReader(os.Stdin)

	name1, err := reader1.ReadString(('\n'))
	s = strings.TrimSpace(name1)
	if err != nil {
		log.Fatal(err)
	}
	ival, err = strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	val2 := ival
	pl("addition of x and y is ", val1+val2)

}
