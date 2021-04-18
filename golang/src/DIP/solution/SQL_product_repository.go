package problem

type SQLProductRepository struct {}

func NewProductRepository() IProductRepository {
	return &SQLProductRepository{}
}

func (s *SQLProductRepository) getAllProductName() []string{
	list := []string {
		"soap","toothpaste",
	}
	return list
}

