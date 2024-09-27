package main

import "fmt"

func main() {
	x := 10
	y := 10

	a := x * y
	b := x / y
	c := x + y
	d := x - y

	fmt.Println(x, "x", y, "=", a)
	fmt.Println(x, "/", y, "=", b)
	fmt.Println(x, "+", y, "=", c)
	fmt.Println(x, "-", y, "=", d)

	//#2 Augmented
	e := 10

	e *= 10
	fmt.Println(e)
	e /= 10
	fmt.Println(e)
	e += 10
	fmt.Println(e)
	e -= 10
	fmt.Println(e)

	//#3 Unary Operator

	f := 10
	g := false

	f++
	fmt.Println(f)
	f--
	fmt.Println(f)
	f = -f
	fmt.Println(f)
	f = +f
	fmt.Println(f)
	g = !g
	fmt.Println(g)

}
