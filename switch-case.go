package main

import "fmt"

func main() {

	//#1 switch case
	player := "ligichi"

	switch player {
	case "ligichi":
		fmt.Println("Hallo", player)
	case "None":
		fmt.Println("Undefind")
	default:
		fmt.Println("Hallo World")

	}

	//#2 Switch Case short condition

	switch job := "Katana"; job == "Katana" {
	case true:
		fmt.Println("Your Job is ", job)
	case false:
		fmt.Println("Unknown Job")
	}

	//#3 Switch Case condition

	job := "Staff"
	switch {
	case job == "Katana":
		fmt.Println("Your Job Is", job)
	case job == "Sword":
		fmt.Println("Your Job Is", job)
	default:
		fmt.Println("Your Job Is", job)
	}
}
