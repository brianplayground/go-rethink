package model

import "time"

type ListItem struct {
	Id string `gorethink:"id,omitempty"`
	Product string
	Quantity int
	Status string
	Created time.Time
	Updated time.Time
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
