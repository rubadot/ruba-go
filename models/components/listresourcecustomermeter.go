package components

type ListResourceCustomerMeter struct {
	Items      []CustomerMeter `json:"items"`
	Pagination Pagination      `json:"pagination"`
}

func (l *ListResourceCustomerMeter) GetItems() []CustomerMeter {
	if l == nil {
		return []CustomerMeter{}
	}
	return l.Items
}

func (l *ListResourceCustomerMeter) GetPagination() Pagination {
	if l == nil {
		return Pagination{}
	}
	return l.Pagination
}
