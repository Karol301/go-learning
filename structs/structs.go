package main

import (
	"fmt"
)

type Car struct {
	Model string
	Price int
	//Color Color
	Color
	// Color struct {
	// 	Type string
	// 	Id   int
	// }
}

type Color struct {
	Type string
	Id   int
}

func (car Car) area() {
	if car.Price >= 100000 {
		fmt.Println("This car is expensive")
	} else {
		fmt.Println("This car is cheap")
	}
}

func main() {
	// myCar := car{}
	// myCar.Model = "Toyota"
	// myCar.Price = 20000
	// myCar.Color = Color{Type: "Red", Id: 1}

	// myCar.Color = struct {
	// 	Type string
	// 	Id   int
	// }{Type: "Red", Id: 1}

	myCar := Car{
		Model: "Toyota",
		Price: 20000,
		Color: Color{
			Type: "Red",
			Id:   1,
		},
	}

	fmt.Println(myCar.Color.Id)
	myCar.area()
}
