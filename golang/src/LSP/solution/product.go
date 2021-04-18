package problem

type Product struct {
	discount float64
}

func NewProduct() *Product {
	return &Product{
		discount: 20.0,
	}
}

func (p *Product) GetDiscount() float64 {
	return p.discount
}

