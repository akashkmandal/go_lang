package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		var number int
		fmt.Print("Please enter a number: ")
		fmt.Scan(&number)
		secret := randomNumber()
		//fmt.Println(secret)
		if secret == number {
			fmt.Println("You guessed the number correct, random number is ", secret, "and your number is ", number)
			break
		} else {
			fmt.Println("You did't guess it correct, Please try again")
		}

	}

}

func randomNumber() int {
	//	min := 0
	//	max := 10
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
