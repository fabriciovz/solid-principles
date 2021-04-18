package problem_lsp

type InHouseProduct struct {
	Product
}
func NewInHouseProduct() *InHouseProduct {
	return &InHouseProduct{
		Product{
			discount: 20.0,
		},
	}
}

func (i *InHouseProduct) applyExtraDiscount(){
	i.discount = i.discount * 1.5
}

