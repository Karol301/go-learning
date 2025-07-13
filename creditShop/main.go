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

func calcCredits(credits int, p Product) int {
	productPrice := p.getPrice()
	if productPrice > credits {
		fmt.Printf("You don't have enough credits!\n")
		return credits
	} else {
		newCredits := credits - productPrice
		fmt.Printf("You bought %s for %d credits. Remaining: %d\n", p.getName(), productPrice, newCredits)
		return newCredits
	}
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

	products := []Product{book, food}

	for _, product := range products {
		getInfo(product)
	}

	fmt.Println("Select a product by name:")

	var userInput string
	fmt.Scanln(&userInput)

	productFound := false
	for _, product := range products {
		if product.getName() == userInput {
			credits = calcCredits(credits, product)
			productFound = true
			break
		}
	}

	if !productFound {
		fmt.Println("No product with that name.")
	}

	fmt.Printf("Remaining credits: %d\n", credits)
}
