package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	milliSeconds := time.Now().Unix()
	rand.Seed(milliSeconds)
	targetNumber := rand.Intn(100) + 1

	fmt.Println("Welcome to Guess the Number!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	fmt.Println("You have 10 attempts to guess the number.")

	var userNumber int 
	var attempts int = 10
	for userNumber != targetNumber && attempts > 0 {
		fmt.Printf("You have %d attempts. Enter your guess: ", attempts)
		fmt.Scanln(&userNumber)


		if userNumber < targetNumber {
			fmt.Println("Your guess is too LOW.")
		} else if userNumber > targetNumber {
			fmt.Println("Your guess is too HIGH.")
		} else {
			fmt.Println("Congratulations! You guessed the number.")
			break
		}
		attempts--
	}

	if attempts == 0 {
		fmt.Println("You've used all your attempts! Guess the number was:", targetNumber)
	}
}