package solution_lsp

import "fmt"

type PricingUtils struct {
}

func Run() {
	p1 := NewProduct()
	p2 := NewProduct()
	p3 := NewInHouseProduct()

	productList := make([]IProduct,0)

	productList = append(productList,p1)
	productList = append(productList,p2)
	productList = append(productList,p3)

	for _,product := range productList {
		//fmt.Println(i)

		fmt.Println(product.GetDiscount())
	}
}
