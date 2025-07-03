package main

import (
	"fmt"
	"math/rand"
)

func sum(x int, y int) int {
	return x + y
}

func random_int() (int, int) {
	num1 := rand.Intn(100)
	num2 := rand.Intn(50)
	return num1, num2
}

func multiplication(x float64, y float64) (mul float64) {
	if y == 0 || x == 0 {
		return mul
	} else {
		mul = x * y
		return mul
	}
}

func main() {
	fmt.Println(sum(1, 2))

	x, y := random_int()
	fmt.Println(x, y)

	z, _ := random_int()
	fmt.Println(z)

	div := multiplication(1, 3)
	fmt.Println(div)
}
