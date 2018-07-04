package repository

import (
	"../model"
	"../config"
	"gopkg.in/gorethink/gorethink.v4"
	"log"
	"time"
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

func (r *ListItemRepository) InsertItem(product string, quantity int) *model.ListItem{
	item := model.NewItem(product,quantity)
	item.Created = time.Now()

	_, err := gorethink.Table("shoppingList").Insert(item).Run(config.RethinkSession())
	if err != nil{
		log.Fatal(err.Error())
	}
	//err = res.All(&item)
	//if err != err { log.Fatal(err.Error()) }
	return item
}