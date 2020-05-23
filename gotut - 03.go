package main

import (
	"fmt"
)

func add(x float64, y float64) float64 {
	return x + y
}

func add2(x, y float64) float64 {
	return x + y
}

func mulitple(a, b string) (string, string) {
	return a, b
}

func main() {
	//var num1 float64 = 5.6
	//var num2 float64 = 9.5
	var num1, num2 float64 = 5.6, 9.5
	fmt.Println(add(num1, num2))
	fmt.Println(add2(num1, num2))

	var a int = 62
	var b float64 = float64(a)
	fmt.Println("b is", b)

	//x := a //x 's type will be int'

	fmt.Println(mulitple("hello", "world"))
}
