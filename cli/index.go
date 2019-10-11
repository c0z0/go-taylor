package main

import (
	"fmt"
	"go-taylor/calculator"
)

func main() {
	var x float64

	fmt.Print("x = ")
	fmt.Scan(&x)

	fmt.Print("e ^ x = ")
	fmt.Println(calculator.Exp(x))
	fmt.Print("ln x = ")
	fmt.Println(calculator.Ln(x))
	fmt.Print("e ^ (-x^2) = ")
	fmt.Println(calculator.Norm(x))
	fmt.Print("sin(x) = ")
	fmt.Println(calculator.Sin(x))
	fmt.Print("cos(x) = ")
	fmt.Println(calculator.Cos(x))
	fmt.Print("tan(x) = ")
	fmt.Println(calculator.Sin(x) / calculator.Cos(x))
	fmt.Print("cot(x) = ")
	fmt.Println(calculator.Cos(x) / calculator.Sin(x))
}
