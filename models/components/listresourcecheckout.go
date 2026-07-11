package components

type ListResourceCheckout struct {
	Items      []Checkout `json:"items"`
	Pagination Pagination `json:"pagination"`
}

func (l *ListResourceCheckout) GetItems() []Checkout {
	if l == nil {
		return []Checkout{}
	}
	return l.Items
}

func (l *ListResourceCheckout) GetPagination() Pagination {
	if l == nil {
		return Pagination{}
	}
	return l.Pagination
}
