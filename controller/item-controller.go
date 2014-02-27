package controller

import (
	"fmt"
	. "github.com/berwyn/masterdex/model"
	"github.com/codegangsta/martini"
	"net/http"
)

type ItemNotFoundError struct {
	slug string
}

func (err ItemNotFoundError) Error() string {
	return fmt.Sprintf("Couldn't find the item with slug %s", err.slug)
}

type ItemValidationError struct{}

func (err ItemValidationError) Error() string {
	return "There was a problem with the item you provided wasn't valid"
}

type ItemController struct {
	datastore Datastore
}

func (ctrl ItemController) Register(server *martini.ClassicMartini) {
	server.Get("/item", ctrl.Index)
	server.Get("/item/:slug", ctrl.Read)
}

func (ctrl ItemController) Index(request *Request) {
	request.Data = new(struct{})
	request.Status = 200
	request.Template = "item"
}

func (ctrl ItemController) Read(params martini.Params, request *Request) {
	item, err := ctrl.datastore.Find(params["slug"])
	if err != nil {
		request.Error(http.StatusNotFound, err.Error())
		return
	}

	request.Data = item
	request.Status = http.StatusOK
	request.Template = "item"
}

func (ctrl ItemController) Create(payload Item, request *Request) {
	item, err := ctrl.datastore.Insert(payload)
	if err != nil {
		if _, ok := err.(*ItemValidationError); ok {
			request.Error(422, err.Error())
		} else {
			request.Error(http.StatusInternalServerError, "We couldn't save the item you provided, please try again later")
		}
		return
	}

	request.Data = item
	request.Status = http.StatusCreated
	request.Template = "item"
}
