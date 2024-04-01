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
	newProd := NewProduct()
	product, err := newProd.GetProductByName(name)
	if err != nil {
		return err
	}
	c.Item = append(c.Item, CartItem{product, quantity})
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
