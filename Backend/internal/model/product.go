package model

type Product struct {
	ProductName string
	Price       string
}

func (p Product) Public() Product {
	return Product{
		ProductName: p.ProductName,
		Price:       p.Price,
	}
}
