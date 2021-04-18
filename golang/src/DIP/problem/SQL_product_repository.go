package problem

type SQLProductRepository struct {}

func NewSQLProductRepository() *SQLProductRepository {
	return &SQLProductRepository{}
}

func (s *SQLProductRepository) getAllProductName() []string{
	list := []string {
		"soap","toothpaste",
	}
	return list
}

