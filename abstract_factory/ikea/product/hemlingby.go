package product

import abstrac_factory "design-pattern/abstract_factory"

type Hemlingby struct {
}

func (h *Hemlingby) Price() float64 {
	return 1795000
}
func (h *Hemlingby) Style() abstrac_factory.SofaStyle {
	return abstrac_factory.EuropeanStyle
}
