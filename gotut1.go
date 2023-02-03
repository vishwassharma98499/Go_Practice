package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"log"
	"os"
)

var pl = f.Println

//
/*
block comment
*/

func main() {

	pl("Hello Go")
	fmt.Println("Hello Go")
	//fmt.Print mean no new line after print

	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}
