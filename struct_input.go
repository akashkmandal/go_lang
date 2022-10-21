package main

import "fmt"

// defining a structure
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

var firstName string
var lastName string
var age int
var salary int

// define an empty slice for storing values based on struct values
//var mySlice = make([]Employee, 0)

func main() {
	fmt.Println("Please enter your First Name:")
	fmt.Scanln(&firstName)
	fmt.Println("Please enter your Last Name:")
	fmt.Scanln(&lastName)
	fmt.Println("Please enter your Age:")
	fmt.Scanln(&age)
	fmt.Println("Please enter you Salary:")
	fmt.Scanln(&salary)
	mySlice := Employee{firstName: firstName, lastName: lastName, age: age, salary: salary}
	fmt.Println(mySlice.firstName)
}
