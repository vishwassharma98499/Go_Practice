package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	iPrimeArray := []int{2, 3, 5, 7, 11}
	var istringarray []string

	for _, val := range iPrimeArray {
		istringarray = append(istringarray, strconv.Itoa(val))
	}

	for _, num := range istringarray {
		_, err = f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		pl("prime : ", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("data.txt")

	if errors.Is(err, os.ErrNotExist) {
		pl("file does not exist")
	} else {
		f, err = os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		if _, err := f.WriteString(" 13 \n"); err != nil {
			log.Fatal(err)
		}
	}

}
