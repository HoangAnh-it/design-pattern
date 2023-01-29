package factory

import "designPattern/creationalDesignPattern/abstractFactory/shoe"

type NikeFactory struct{}

func (nf *NikeFactory) MakeShoe() shoe.IShoe {
	nikeShoe := shoe.NikeShoe{}
	nikeShoe.SetBrand("nike")
	nikeShoe.SetSize(99)
	return &nikeShoe
}
