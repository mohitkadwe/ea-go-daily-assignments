package main

import (
	"flag"
	"fmt"
	"math"
)

type calculator struct {
	x int
	y int
}

func (c *calculator) add() int {
	return c.x + c.y
}

func (c *calculator) sub() int {
	return c.x - c.y
}

func (c *calculator) mul() int {
	return c.x * c.y
}

func (c *calculator) div() int {
	return c.x / c.y
}

func (c *calculator) trigonometricSin() float64 {
	result := math.Sin(float64(c.x))
	return result
}

func (c *calculator) trigonometricCos() float64 {
	result := math.Cos(float64(c.x))
	return result
}

func (c *calculator) trigonometricTan() float64 {
	result := math.Tan(float64(c.x))
	return result
}

func (c *calculator) squareRoot() float64 {
	result := math.Sqrt(float64(c.x))
	return result
}

func main() {

	add := flag.Bool("add", false, "Add two numbers")
	sub := flag.Bool("sub", false, "Subtract two numbers")
	mul := flag.Bool("mul", false, "Multiply two numbers")
	div := flag.Bool("div", false, "Divide two numbers")
	sin := flag.Bool("sin", false, "Trigonometric Sin")
	cos := flag.Bool("cos", false, "Trigonometric Cos")
	tan := flag.Bool("tan", false, "Trigonometric Tan")
	sqrt := flag.Bool("sqrt", false, "Square Root")

	flag.Parse()

	a := 0
	b := 0

	switch {
	case *add:
		fmt.Println("Enter first number")
		fmt.Scan(&a)
		fmt.Println("Enter second number")
		fmt.Scan(&b)
		c := calculator{a, b}
		fmt.Println(c.add())
	case *sub:
		fmt.Println("Enter first number")
		fmt.Scan(&a)
		fmt.Println("Enter second number")
		fmt.Scan(&b)
		c := calculator{a, b}
		fmt.Println(c.sub())
	case *mul:
		fmt.Println("Enter first number")
		fmt.Scan(&a)
		fmt.Println("Enter second number")
		fmt.Scan(&b)
		c := calculator{a, b}
		fmt.Println(c.mul())
	case *div:
		fmt.Println("Enter first number")
		fmt.Scan(&a)
		fmt.Println("Enter second number")
		fmt.Scan(&b)
		c := calculator{a, b}
		fmt.Println(c.div())
	case *sin:
		fmt.Println("Enter a number")
		fmt.Scan(&a)
		c := calculator{a, 0}
		fmt.Println(c.trigonometricSin())
	case *cos:
		fmt.Println("Enter a number")
		fmt.Scan(&a)
		c := calculator{a, 0}
		fmt.Println(c.trigonometricCos())
	case *tan:
		fmt.Println("Enter a number")
		fmt.Scan(&a)
		c := calculator{a, 0}
		fmt.Println(c.trigonometricTan())
	case *sqrt:
		fmt.Println("Enter a number")
		fmt.Scan(&a)
		c := calculator{a, 0}
		fmt.Println(c.squareRoot())
	}

}
