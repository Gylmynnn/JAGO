package main

import "fmt"

func main() {

	//#1 if-else
	player := "ligichi"

	if player == "ligichi" {
		fmt.Println("Player :", player)
	} else if player == "gyl" {
		fmt.Println("Player :", player)
	} else {
		fmt.Println("Undefind")
	}

	//#2 if short

	if length := len(player); length < 4 {
		fmt.Println("Username to short")
	} else {
		fmt.Println("Welkam", player)
	}
}
