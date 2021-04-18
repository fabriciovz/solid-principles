package problem_lsp

type Product struct {
	discount float64
}

type IProduct interface {

}

func NewProduct() *Product {
	return &Product{
		discount: 20.0,
	}
}

func (p *Product) GetDiscount() float64 {
	return p.discount
}

