package controllers

import (
	"encoding/json"

	"github.com/casbin/casibase/object"
)

func (c *ApiController) GetGlobalStores() {
	c.ResponseOk(object.GetGlobalStores())
}

func (c *ApiController) GetStores() {
	owner := c.Input().Get("owner")

	c.ResponseOk(object.GetStores(owner))
}

func (c *ApiController) GetStore() {
	id := c.Input().Get("id")

	store := object.GetStore(id)
	if store == nil {
		c.ResponseError("store is empty")
		return
	}

	store.Populate()

	c.ResponseOk(store)
}

func (c *ApiController) UpdateStore() {
	id := c.Input().Get("id")

	var store object.Store
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &store)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.UpdateStore(id, &store))
}

func (c *ApiController) AddStore() {
	var store object.Store
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &store)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.AddStore(&store))
}

func (c *ApiController) DeleteStore() {
	var store object.Store
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &store)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.DeleteStore(&store))
}
