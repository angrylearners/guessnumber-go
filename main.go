package main

import (
	"fmt"
	"math/rand"
)

func inputOneNumber() (res int) {
retry:
	if _, err := fmt.Scanf("%d\n", &res); err != nil {
		goto retry
	}
	return
}

func inputRange() (from, to int) {
retry:
	fmt.Print("Input the range of the secret number first. From= ")
	from = inputOneNumber()
	fmt.Print("To= ")
	to = inputOneNumber()

	if from >= to {
		fmt.Printf("Invalid range [%d,%d).\n", from, to)
		goto retry
	}
	return
}

func generateSecretNumber(from, to int) int {
	fmt.Printf("Secret number will be generated in range [%d,%d).\n", from, to)
	return rand.Intn(to-from) + from
}

func guessNumber(num int) {
	for {
		fmt.Print("Guess what the number is: ")
		guess := inputOneNumber()
		switch {
		case guess > num:
			fmt.Printf("%d is too big.\n", guess)
		case guess < num:
			fmt.Printf("%d is too small.\n", guess)
		case guess == num:
			fmt.Printf("Congratulation! %d is the secret number\n", guess)
			return
		}
	}
}

func toPlayAgain() bool {
retry:
	fmt.Print("Do you want to play again? [y/n] ")
	var res string
	if _, err := fmt.Scanf("%s\n", &res); err != nil {
		fmt.Println("Please input y or n.")
		goto retry
	}
	switch res {
	case "y":
		return true
	case "n":
		return false
	default:
		fmt.Println("Please input y or n.")
		goto retry
	}
}

func main() {
	for {
		from, to := inputRange()
		fmt.Printf("The secret number will be in range [%d,%d).\n", from, to)
		secretNumber := generateSecretNumber(from, to)
		guessNumber(secretNumber)
		if !toPlayAgain() {
			return
		}
	}
}
