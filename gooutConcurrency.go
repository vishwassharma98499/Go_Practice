package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("func 1: ", i)
	}
}

func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("func 2: ", i)
	}
}

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3

}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6

}

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance

}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("not enuf money in account")
	} else {
		fmt.Printf("%d withdrawn balance %d\n", v, a.balance)
		a.balance -= v
	}

}

// concurrency is basically allow us to run multiple blocks of code share execution time by pausing executionand we can also run bloacks of code in parallel time in go concurrency task known as go routine
func main() {
	//go printTo10()
	//go printTo15()
	//time.Sleep(2 * time.Second)
	// when main end entire programe ends so pause the main for 2 seconds

	//channel1 := make(chan int)
	//channel2 := make(chan int)

	//go nums1(channel1)
	//go nums2(channel2)
	//pl(<-channel1)
	//pl(<-channel2)
	//pl(<-channel1)
	//pl(<-channel2)
	//pl(<-channel1)
	//pl(<-channel2)

	var acct Account
	acct.balance = 100

	pl("balance :", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)

	}

	time.Sleep(2 * time.Second)
}
