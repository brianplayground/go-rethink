package controller

import (
	"../model"
	"../repository"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Controller struct {
	Repository repository.ListItemRepository
} 

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	items := c.Repository.GetItems()
	data, _ := json.Marshal(items)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) FindById(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id := params["id"]
	item := c.Repository.GetItemById(id)
	data, _ := json.Marshal(item)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return

}

func (c *Controller) Insert(w http.ResponseWriter, r *http.Request){
	var i model.ListItem
	if r.Body == nil{
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	c.Repository.InsertItem(i.Product,i.Quantity)
	res, _ := json.Marshal(i)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request){
	var i model.ListItem
	params := mux.Vars(r)
	id := params["id"]

	if r.Body == nil{
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	c.Repository.UpdateItem(id,i)
	res, _ := json.Marshal(i)
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	if id == ""{
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	c.Repository.DeleteItem(id)
	w.WriteHeader(http.StatusOK)
}