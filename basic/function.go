package main

import "fmt"

func main() {

	result := getHallo("Hallow")
	fmt.Println(result)

	sayHallo()

	withParam("Muhamad", "Gilman")

}

//#1

func sayHallo() {
	fmt.Println("Hallo World")
}

//#2

func withParam(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

//#3

func getHallo(text string) string {
	return text
}
