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
	pl("how old are you")
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
	pl(ival)
}
