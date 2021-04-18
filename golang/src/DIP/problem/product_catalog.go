package problem

import "fmt"

type ProductCatalog struct {}

func (s *ProductCatalog) listAllProducts() {
	sqlProductRepository:= NewSQLProductRepository()
	list := sqlProductRepository.getAllProductName()
	fmt.Println(list)
}