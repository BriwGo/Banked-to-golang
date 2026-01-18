package main

import (
	"bank/app"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	NewUser := app.NewStorage("Alice", 200)

	for {

		balance, err := NewUser.GetBalance()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Current balance: %d\n", balance)
		time.Sleep(200 * time.Millisecond)
		err = NewUser.Pay(rand.Intn(100))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(200 * time.Millisecond)

		err = NewUser.Deposit(rand.Intn(50))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(200 * time.Millisecond)

	}
}
