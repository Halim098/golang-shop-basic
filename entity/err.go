package entity

type CustomError struct {
	MsgError string
}

func (e CustomError) Error() string {
	return e.MsgError
}
