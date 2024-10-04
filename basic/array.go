package main

import "fmt"

func main() {

	//#1

	var names [3]string

	names[0] = "Muhamad"
	names[1] = "Gilman"
	names[2] = "Nurrajabi"

	fmt.Println(names[0], names[1], names[2])

	//#2

	values := [4]int{
		85, 90, 95, 99,
	}

	fmt.Println(values)
  fmt.Println(len(values))
  fmt.Println(values[3])
  values[3] = 100
  fmt.Println(values)

}
