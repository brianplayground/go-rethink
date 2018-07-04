package model

import "time"

type ListItem struct {
	Id string `gorethink:"id,omitempty"`
	Product string `gorethink:"product"`
	Quantity int `gorethink:"quantity"`
	Status string `gorethink:"status"`
	Created time.Time `gorethink:"created_at"`
	Updated time.Time `gorethink:"updated_at"`
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
