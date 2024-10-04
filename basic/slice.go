package main

import "fmt"

func main() {

	//#1 slice

	names := []string{"Muhamad", "Gilman", "Nurrajabi", "Gylmyn", "De", "Rune"}

	slice := names[:]
	slice2 := names[1:4]

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice2)
	fmt.Println(slice2[0])

	fmt.Println(len(slice))
	fmt.Println(len(slice2))

	//#2 append slice

	fruits := []string{"Banana", "Mango", "Apple", "Dragon"}
	fruitSlice1 := fruits[2:]
	fruitSlice1[0] = "Strawberry"
	fruitSlice1[1] = "Melon"
	fmt.Println(fruits)

	fruitSlice2 := append(fruitSlice1, "Rasberry")
	fruitSlice2[0] = "Fruits"
	fmt.Println(fruitSlice2)
	fmt.Println(fruits)

	//#3 make slice
	// []type, length, capacity
	newSlice := make([]string, 3, 5)
	newSlice[0] = "Muhamad"
	newSlice[1] = "Gilman"
	newSlice[2] = "Nurrajabi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//#4 copy slice
	fromFruitsSlice := fruits[:]
	toNewSlice := make([]string, len(fromFruitsSlice), cap(fromFruitsSlice))

	copy(toNewSlice, fromFruitsSlice)
	fmt.Println(toNewSlice)

}

