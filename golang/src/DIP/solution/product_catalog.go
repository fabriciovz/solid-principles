package problem

import "fmt"

type ProductCatalog struct {}

func (s *ProductCatalog) listAllProducts() {
	productRepository:= NewProductRepository()
	list := productRepository.getAllProductName()
	fmt.Println(list)
}