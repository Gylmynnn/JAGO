package main

import (
	"fmt"

)

func main()  {
  type Name string
  type Age int16

  var name Name = "Muhamad Gilman"
  fmt.Println(name)
  fmt.Println(Name("Nurrajabi"))

  var age Age = 20
  fmt.Println(age)
  fmt.Println(Age(2004))
}
