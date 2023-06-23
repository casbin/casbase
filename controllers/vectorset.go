package controllers

import (
	"encoding/json"

	"github.com/casbin/casibase/object"
)

func (c *ApiController) GetGlobalVectorsets() {
	c.ResponseOk(object.GetGlobalVectorsets())
}

func (c *ApiController) GetVectorsets() {
	owner := c.Input().Get("owner")

	c.ResponseOk(object.GetVectorsets(owner))
}

func (c *ApiController) GetVectorset() {
	id := c.Input().Get("id")

	c.ResponseOk(object.GetVectorset(id))
}

func (c *ApiController) UpdateVectorset() {
	id := c.Input().Get("id")

	var vectorset object.Vectorset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &vectorset)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.UpdateVectorset(id, &vectorset))
}

func (c *ApiController) AddVectorset() {
	var vectorset object.Vectorset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &vectorset)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.AddVectorset(&vectorset))
}

func (c *ApiController) DeleteVectorset() {
	var vectorset object.Vectorset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &vectorset)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.DeleteVectorset(&vectorset))
}
