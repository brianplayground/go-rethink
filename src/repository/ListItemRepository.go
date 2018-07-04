package repository

import (
	"../model"
	"../config"
	"gopkg.in/gorethink/gorethink.v4"
	"log"
	"time"
)
type ListItemRepository struct{}

var tableName = "shoppingList"

func (r *ListItemRepository) GetItems() model.ListItems{
	items := model.ListItems{}

	res, err := gorethink.Table(tableName).Run(config.RethinkSession())
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

	res , err := gorethink.Table(tableName).Insert(item).Run(config.RethinkSession())
	if err != nil{
		log.Fatal(err.Error())
	}
	defer res.Close()
	return item
}

func (r *ListItemRepository) UpdateItem(id, product string, quantity int)  {

}

func (r *ListItemRepository) GetItemById(id string) []model.ListItem{
	item := []model.ListItem{}
	res, err := gorethink.Table(tableName).Get(id).Run(config.RethinkSession())
	if err != nil{
		log.Fatal(err.Error())
	}
	err = res.All(&item)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Close()

	return item
}