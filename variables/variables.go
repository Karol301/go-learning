package main

import "fmt"

func main() {
	//int, uint, float, string | float32, float64 ...
	x := 1.6 //var x float = 1
	y := int(x)
	fmt.Println(x, y)

	if x < 2 {
		fmt.Println("x is less than 2")
	} else if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is between 2 and 10")
	}
}
