package repository

import (
	"../model"
)
type ListItemRepository struct{}

func (r *ListItemRepository) GetItems() model.ListItems{
	return model.ListItems{
		model.ListItem{
			"1",
			"Milk",
			1,
			"false",
		},
		model.ListItem{
			"2",
			"Something",
			4,
			"false",
		},
	}
}
