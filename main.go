package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a secret number
	secret := getRandomNumber()
	//fmt.Println(secret)
	for matching := false; !matching; {
		// Get user input
		guess := getUserInput()
		//fmt.Println(secret, guess)

		// Make comparison (secret vs guess)
		matching = isMatching(secret, guess)
		fmt.Println(matching)
	}
}
func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Your Guess is too small")
		return false
	} else {
		fmt.Println("YOU GOT IT!")
		return true
	}
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

func getUserInput() int {
	fmt.Println("Enter input your guess:")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}
	return input
}
