# Guess Number [![CircleCI](https://circleci.com/gh/angrylearners/guessnumber-go.svg?style=svg)](https://circleci.com/gh/angrylearners/guessnumber-go)
An alternative implementation of number guessing game, in go. 

> The Goal:The program will first randomly generate a number unknown to the user. The user needs to guess what that number is. (In other words, the user needs to be able to input information.) If the user’s guess is wrong, the program should return some sort of indication as to how wrong (e.g. The number is too high or too low). If the user guesses correctly, a positive indication should appear. You’ll need functions to check if the user input is an actual number, to see the difference between the inputted number and the randomly generated numbers, and to then compare the numbers.

## How to play
Compile the program with:
```shell script
go get github.com/angrylearners/guessnumber-go
```
Then you can find the executable file named `guessnumber-go` in `$GOPATH/bin`.
Open `guessnumber-go` in your shell:
```
Welcome to number guessing game.
Input the range of the secret number first. From= 0
To= 100
The secret number will be in range [0,100).
Secret number will be generated in range [0,100).
Guess what the number is: 50
50 is too small.
Guess what the number is: 75
75 is too small.
Guess what the number is: 81
Congratulation! 81 is the secret number
Do you want to play again? [y/n] n
Goodbye.
```
