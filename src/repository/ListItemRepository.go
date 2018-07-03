package repository

import (
	"../model"
	"../config"
	"gopkg.in/gorethink/gorethink.v4"
	"log"
)
type ListItemRepository struct{}

func (r *ListItemRepository) GetItems() model.ListItems{
	items := model.ListItems{}

	res, err := gorethink.Table("shoppingList").Run(config.RethinkSession())
	if err != nil {
		log.Fatal(err.Error())
	}
	err = res.All(&items)
	if err != nil{
		log.Fatal(err.Error())
	}

	defer res.Close()
	return items
}

