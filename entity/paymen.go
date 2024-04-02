package entity

type Payment struct {
	Item   []CartItem
	Total  int
	Money  int
	Change int
}

func NewPayment() *Payment {
	return &Payment{}
}

func (p *Payment) AddItem(c *Cart) error {
	if c == nil {
		return CustomError{"Cart is empty"}
	}
	p.Item = c.Item
	for _, v := range p.Item {
		p.Total += v.Product.Price * v.Quantity
	}
	return nil
}

func (p *Payment) Pay(money int) (int, string, error) {
	if money < p.Total {
		return 0, "", CustomError{"Money is not enough"}
	}
	p.Money = money
	p.Change = money - p.Total
	return p.Change, "Success pay", nil
}
