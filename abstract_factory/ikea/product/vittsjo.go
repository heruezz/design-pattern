package product

import abstrac_factory "design-pattern/abstract_factory"

type Vittsjo struct {
}

func (v *Vittsjo) Price() float64 {
	return 899000
}

func (v *Vittsjo) Size() abstrac_factory.Dimension {
	return abstrac_factory.Dimension{
		Length: 80,
		Width:  78,
		Height: 40,
	}
}

func (v *Vittsjo) IsFoldable() bool {
	return false
}
