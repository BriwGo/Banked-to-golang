package app

import (
	"fmt"
	"math/rand"
)

type Storage struct {
	ID      int
	Name    string
	Balance int
}

func NewStorage(name string, balance int) *Storage {
	return &Storage{
		ID:      rand.Int(),
		Name:    name,
		Balance: balance,
	}
}

func (l *Storage) GetBalance() (int, error) {
	return l.Balance, nil
}

func (l *Storage) Pay(amount int) error {
	randomValue := rand.Intn(100)

	if randomValue < 30 {
		return NetworkError
	}

	if amount > l.Balance {
		return insufficientFunds
	}

	l.Balance -= amount
	fmt.Println("Баланс упал,транзакция упала")
	return nil
}

func (l *Storage) Deposit(amount int) error {
	if amount < 0 {
		return WhatTheFuck
	}
	fmt.Println("Баланс пополнен на ", amount)
	l.Balance += amount
	return nil
}
