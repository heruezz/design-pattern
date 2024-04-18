package ikea

import (
	abstrac_factory "design-pattern/abstract_factory"
	"design-pattern/abstract_factory/informa/product"
)

type Informa struct {
}

func (i *Informa) CreateChair() abstrac_factory.Chair {
	// product := db.Query()

	return &product.BeanBag{
		// price:        product.Price,
		// SofnessLevel: product.SoftnessLevel,
	}
}

func (i *Informa) CreateCofeeTable() abstrac_factory.CoffeeTable {
	return nil
}

func (i *Informa) CreateSofa() abstrac_factory.Sofa {
	return nil
}
