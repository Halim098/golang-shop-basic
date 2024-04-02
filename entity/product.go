package entity

type Product struct {
	Name  string
	Price int
	Stock int
}

var DataProduct = []Product{
	{Name: "Book", Price: 5000, Stock: 10},
	{Name: "Pencil", Price: 2000, Stock: 10},
	{Name: "Pen", Price: 4000, Stock: 10},
	{Name: "Eraser", Price: 3000, Stock: 10},
}

func NewProduct() *Product {
	return &Product{}
}

func (dp *Product) ReduceStock(name string, quantity int) error {
	for i := range DataProduct {
		if name == DataProduct[i].Name {
			if DataProduct[i].Stock < quantity {
				return CustomError{"Stock is not enough"}
			}
			DataProduct[i].Stock -= quantity
			return nil
		}
	}
	return CustomError{"Product not found"}
}

func (dp *Product) GetProductByName(name string) (Product, error) {
	for i := range DataProduct {
		if DataProduct[i].Name == name {
			return DataProduct[i], nil
		}
	}
	return Product{}, CustomError{"Product not found"}
}
