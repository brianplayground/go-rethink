package controller

import (
	"../model"
	"../repository"
	"net/http"
	"encoding/json"
)

type Controller struct {
	Repository repository.ListItemRepository
} 

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	items := c.Repository.GetItems()
	data, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) Insert(w http.ResponseWriter, r *http.Request){
	var i model.ListItem
	if r.Body == nil{
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	c.Repository.InsertItem(i.Product,i.Quantity)
	w.WriteHeader(http.StatusCreated)
	res, _ := json.Marshal(i)
	w.Write(res)
}