package factory

import "designPattern/creationalDesignPattern/abstractFactory/shoe"

type IFactory interface {
	MakeShoe() shoe.IShoe
}

func GetFactory(factory string) IFactory {
	switch factory {
	case "adidas":
		return &AdidasFactory{}
	case "nike":
		return &NikeFactory{}
	default:
		return nil
	}
}
