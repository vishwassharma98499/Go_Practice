package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	loc, err := time.LoadLocation(("America/New_york"))
	if err != nil {
		pl(err)
	}
	fmt.Printf("time in newyork is %s\n", now.In(loc))
	loc, _ = time.LoadLocation(("Asia/Shanghai"))
	if err != nil {
		pl(err)
	}
	fmt.Printf("time in Shanghai is %s\n", now.In(loc))
}
