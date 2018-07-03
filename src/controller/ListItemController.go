package controller

import (
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
