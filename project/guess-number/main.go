package main

import (
	"fmt"
	"math/rand"
)

func main() {

	target := int(rand.Intn(100) + 1)

	var guess int
	attemps := 0

	for {
		fmt.Print("Guess a number between 1 and 100")
		fmt.Scanf("%d", &guess)
		attemps++

		if guess < target {
			fmt.Println("Too Low")
		} else if guess > target {
			fmt.Println("Too High")
		} else {
			break
		}

		if attemps == 10 {
			break
		}

		if attemps != max(attemps) && attemps != 10 {
			continue
		} else {
			break
		}

	}
	fmt.Printf("\nTarget was %v ", target)
	fmt.Printf("\nYou Guessed it in %v ", attemps)

}
