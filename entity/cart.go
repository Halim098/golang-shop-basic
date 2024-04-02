package entity

type CartItem struct {
	Product  Product
	Quantity int
}

type Cart struct {
	Item []CartItem
}

func NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) AddItem(name string, quantity int) error {
	if quantity <= 0 {
		return CustomError{"Quantity must be greater than 0"}
	}

	newProd := NewProduct()
	product, err := newProd.GetProductByName(name)
	if err != nil {
		return err
	}
	c.Item = append(c.Item, CartItem{product, quantity})
	err = newProd.ReduceStock(name, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cart) GetCart() []CartItem {
	return c.Item
}

func (c *Cart) RemoveItem(name string) (string, error) {
	for i, v := range c.Item {
		if v.Product.Name == name {
			c.Item = append(c.Item[:i], c.Item[i+1:]...)
			return "Success remove from cart", nil
		}
	}
	return "", CustomError{"Product not found"}
}

func (c *Cart) ResetCart() {
	c.Item = []CartItem{}
}
