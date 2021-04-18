package problem_lsp

type PricingUtils struct {
}

func Run() {
	p1 := NewProduct()
	p2 := NewProduct()
	//p3 := NewInHouseProduct()

	productList := make([]*Product,3)

	productList = append(productList,p1)
	productList = append(productList,p2)
	//productList = append(productList,p3)

	//for (Product product: productList) {
	//
	//	if(product instanceof InHouseProduct) {
	//		((InHouseProduct) product).applyExtraDiscount();
	//	}
	//	System.out.println(product.getDiscount());

}
