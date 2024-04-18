package product

type BeanBag struct {
	// price        float64
	// SofnessLevel int
}

func (b BeanBag) Price() float64 {
	return 1199000
	// return b.price
}

func (b BeanBag) IsIoTEnabled() bool {
	return false
}

func (b BeanBag) IsSoft() bool {
	return true
	// return b.SofnessLevel > 10
}
