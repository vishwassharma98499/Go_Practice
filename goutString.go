package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	sV1 := "A Word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length :", len(sV2))
	pl("Contains :ANother", strings.Contains(sV2, "Another"))
	pl("o Index :", strings.Index(sV2, "o"))
	pl("Replace : ", strings.Replace(sV2, "o", "0", 2)) //2 mean first 2, -1 means all of matches
	sV3 := "\nSpme Words \n"
	sV3 = strings.TrimSpace(sV3)
	pl("Splits :", strings.Split("a-b-c-d", "-"))
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))
	pl("prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

	rStr := "abcdefg"
	pl("Rune cound : ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)

	}

}
