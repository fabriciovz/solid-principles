package solution_lsp

type InHouseProduct struct {
	*Product
}
func NewInHouseProduct() IProduct {
	return &InHouseProduct{
		&Product{
			discount: 20.0,
		},
	}
}

func (i *InHouseProduct) ApplyExtraDiscount(){
	i.discount = i.discount * 1.5
}
func (p *InHouseProduct) GetDiscount() float64 {
	p.ApplyExtraDiscount()
	return p.discount
}

