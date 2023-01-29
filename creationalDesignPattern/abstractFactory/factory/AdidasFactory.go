package factory

import "designPattern/creationalDesignPattern/abstractFactory/shoe"

type AdidasFactory struct{}

func (af *AdidasFactory) MakeShoe() shoe.IShoe {
	shoeAdidas := shoe.AdidasShoe{}
	shoeAdidas.SetBrand("adidas")
	shoeAdidas.SetSize(10)
	return &shoeAdidas
}
