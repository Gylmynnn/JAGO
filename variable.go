package main

import "fmt"

func main() {

	//#1 var late initialisation
	var name string

	name = "Muhamad Gilman"
	fmt.Println("My First Name Is :", name)

	name = "Gilman Nurrajabi"
	fmt.Println("My Last Name Is :", name)

	//#2 var first initialisation
	var name2 = "Muhamad Gilman"
	fmt.Println("My First Name Is :", name2)

	name2 = "Gilman Nurrajabi"
	fmt.Println("My Last Name Is :", name2)

	//#3 initialisation without var
	name3 := "Muhamad Gilman"
	fmt.Println("My First Name Is :", name3)

	name3 = "Gilman Nurrajabi"
	fmt.Println("My Last Name Is :", name3)

	//#4 multiple var variable

	var (
		firstName  = "Muhamad"
		middleName = "Gilman"
		lastName   = "Nurrajabi"
	)

	fmt.Println("My First Name =", firstName)
	fmt.Println("My Middle Name =", middleName)
	fmt.Println("My Last Name =", lastName)

	//#5 const variable (can't change value)

	const name4 string = "Muhamad"
	const name5 = "Gilman"

	fmt.Println("My First Name =", name4)
	fmt.Println("My Middle Name =", name5)

	//#6 multiple const variable

	const (
		firstName2  = "Muhamad"
		middleName2 = "Gilman"
		lastName2   = "Nurrajabi"
	)

	fmt.Println("My First Name =", firstName2)
	fmt.Println("My Middle Name =", middleName2)
	fmt.Println("My Last Name =", lastName2)

}
