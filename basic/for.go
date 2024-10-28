package main

import (
	"fmt"
)

func main() {

	//#1
	counter := 1
	for counter < 10 {
		fmt.Println("#")
		counter++
	}

  // for p := 1; p <= 2; p++ {
  //   fmt.Print(p)
  // }

  //#2
	for i := 1; i <= 3; i++ {
		for g := 1; g <= i; g++ {
			fmt.Print(g)
		}
     fmt.Println("")
	}

  //#3

  jobs := []string{"Katana", "Staff", "Knuckles", "Sword", "Bow"}

  for index, job := range jobs {
    fmt.Println(index, job) 
  }

}
