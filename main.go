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

func main() {
inputRange:
	fmt.Print("Input the range of the secret number first. From= ")
	from := inputOneNumber()
	fmt.Print("To= ")
	to := inputOneNumber()

	if from >= to {
		fmt.Printf("Invalid range [%d,%d)", from, to)
		goto inputRange
	}

	fmt.Printf("The secret number will be in range [%d,%d).\n", from, to)

	num := rand.Intn(to-from) + from

	func() {
		for {
			fmt.Print("Guess what the number is: ")
			guess := inputOneNumber()
			switch {
			case guess > num:
				fmt.Printf("%d is too big\n", guess)
			case guess < num:
				fmt.Printf("%d is too small\n", guess)
			case guess == num:
				fmt.Printf("Congratulation! %d is the secret number\n", guess)
				return
			}
		}
	}()

toContinueOrNot:
	fmt.Print("Do you want to continue? [y/n] ")
	var res string
	if _, err := fmt.Scanf("%s\n", &res); err != nil {
		fmt.Println("Please input y or n.")
		goto toContinueOrNot
	}
	switch res {
	case "y":
		goto inputRange
	case "n":
		return
	default:
		fmt.Println("Please input y or n.")
		goto toContinueOrNot
	}

}
