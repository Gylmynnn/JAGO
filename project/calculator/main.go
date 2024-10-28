package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Input first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Input an operator (*,/,+,-): ")
	fmt.Scanln(&operator)

	fmt.Print("Input second number: ")
	fmt.Scanln(&num2)

	result := calculate(num1, operator, num2)

	fmt.Printf("Result: %f\n", result)
}

func calculate(num1 float64, operator string, num2 float64) float64 {
	result := 0.0

	switch operator {
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Can't divide by zero!")
		}
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	default:
		fmt.Println("Operator Invalid")
	}
	return result
}
