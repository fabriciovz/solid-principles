package solution_lsp

type Product struct {
	discount float64
}

type IProduct interface {
	GetDiscount() float64
}

func NewProduct() IProduct {
	return &Product{
		discount: 20.0,
	}
}

func (p *Product) GetDiscount() float64 {
	return p.discount
}

