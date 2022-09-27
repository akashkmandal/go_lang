package main

import "fmt"

var Number1 int
var Number2 int
var Choice string

func main() {
	for {
		userInput()

		if Choice == "1" || Choice == "Add" {
			fmt.Println(addition(Number1, Number2))
			break
		} else if Choice == "2" || Choice == "multi" {
			fmt.Println(multiply(Number1, Number2))
			break
		} else if Choice == "3" || Choice == "substract" {
			fmt.Println(substraction(Number1, Number2))
			break
		} else if Choice == "4" || Choice == "divide" {
			fmt.Println(division(Number1, Number2))
			break
		} else {
			fmt.Println("Not a valid choice, Please try again")
		}

	}

}

func userInput() {

	fmt.Println("This is a basic calculator written in golang")
	fmt.Print("Please enter the choice between, 1. Add, 2. multi, 3. substract, 4. divide: ")
	fmt.Scan(&Choice)
	fmt.Print("Please enter a number: ")
	fmt.Scan(&Number1)
	fmt.Print("Please enter second number: ")
	fmt.Scan(&Number2)
}

func multiply(Number1 int, Number2 int) int {
	result := Number1 * Number2
	return result
}

func addition(Number1 int, Number2 int) int {
	result := Number1 + Number2
	return result
}

func substraction(Number1 int, Number2 int) int {
	result := Number1 - Number2
	return result
}

func division(Number1 int, Number2 int) int {
	result := Number1 / Number2
	return result
}
