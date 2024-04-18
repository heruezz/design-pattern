package ikea

import (
	abstrac_factory "design-pattern/abstract_factory"
	"design-pattern/abstract_factory/ikea/product"
)

type Ikea struct {
}

func (i *Ikea) CreateChair() abstrac_factory.Chair {
	return &product.Leifarne{}
}

func (i *Ikea) CreateCofeeTable() abstrac_factory.CoffeeTable {
	return &product.Vittsjo{}
}

func (i *Ikea) CreateSofa() abstrac_factory.Sofa {
	return &product.Hemlingby{}
}
