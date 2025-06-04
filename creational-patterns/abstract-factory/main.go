package main

import (
	"fmt"
)

func main() {
	addidasFactory, _ := GetISportFactory("addidas")
	nikeFactory, _ := GetISportFactory("nike")

	addidasShoe := addidasFactory.makeShoe()
	addidasShirt := addidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	printShoeDetails(addidasShoe)
	printShoeDetails(nikeShoe)

	printShirtDetails(addidasShirt)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
