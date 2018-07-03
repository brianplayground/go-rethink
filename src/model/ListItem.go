package model

type ListItem struct {
	ID string
	Product string
	Quantity int
	Status string
}

func (I *ListItem) Bought() bool {
	return I.Status == "true"
}

func NewItem(product string, quantity int) *ListItem{
	return &ListItem{
		Product:product,
		Quantity:quantity,
		Status:"false",
	}
}

type ListItems []ListItem
