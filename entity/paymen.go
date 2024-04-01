package entity

type Payment struct {
	Item  []CartItem
	Total int
	Money int
	Change int
}

func NewPayment() *Payment {
	return &Payment{}
}


