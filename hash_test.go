package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestGenerateFromPassword(t *testing.T) {
	password := "135798642mrongoz"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	fmt.Println("hashedPassword =>", string(hashedPassword))

	cost, err := bcrypt.Cost(hashedPassword)

	if err != nil {
		panic(err)
	}

	fmt.Println("Cost(hashedPassword) =>", cost)

	fmt.Println("password =>", password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	
	if err != nil {
		panic(err)
	}
}
