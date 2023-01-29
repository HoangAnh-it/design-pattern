package main

import (
	factory "designPattern/creationalDesignPattern/abstractFactory/factory"
	"fmt"
)

func main() {
	adidasFactory := factory.GetFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	fmt.Println(adidasShoe)

	nikeFactory := factory.GetFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	fmt.Println(nikeShoe)
}
