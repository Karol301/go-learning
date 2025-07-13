package main

import "fmt"

type Product interface {
	getPrice() int
	getName() string
}

type Book struct {
	Name   string
	Price  int
	Author Author
}

type Author struct {
	Name string
	Age  int
}

type Food struct {
	Name  string
	Price int
}

func (b Book) getPrice() int {
	return b.Price
}

func (b Book) getName() string {
	return b.Name
}

func (f Food) getPrice() int {
	return f.Price
}

func (f Food) getName() string {
	return f.Name
}

func getInfo(p Product) {
	fmt.Printf("Product Name: %s, Price: %d\n", p.getName(), p.getPrice())
}

func main() {
	credits := 100
	fmt.Printf("You have %d credits\n", credits)

	book := Book{
		Name:  "XYZ",
		Price: 10,
		Author: Author{
			Name: "Jan Kowalski",
			Age:  50,
		},
	}

	food := Food{
		Name:  "Pizza",
		Price: 15,
	}

	fmt.Printf("Select a product by name")
	getInfo(book)
	getInfo(food)
	var userInput string
	fmt.Scanln(&userInput)
}
