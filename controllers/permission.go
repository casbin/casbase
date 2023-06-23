package controllers

import (
	"encoding/json"

	"github.com/casbin/casibase/casdoor"
)

func (c *ApiController) GetPermissions() {
	owner := c.Input().Get("owner")

	c.ResponseOk(casdoor.GetPermissions(owner))
}

func (c *ApiController) GetPermission() {
	id := c.Input().Get("id")

	c.ResponseOk(casdoor.GetPermission(id))
}

func (c *ApiController) UpdatePermission() {
	id := c.Input().Get("id")

	var permission casdoor.Permission
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permission)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(casdoor.UpdatePermission(id, &permission))
}

func (c *ApiController) AddPermission() {
	var permission casdoor.Permission
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permission)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(casdoor.AddPermission(&permission))
}

func (c *ApiController) DeletePermission() {
	var permission casdoor.Permission
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &permission)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(casdoor.DeletePermission(&permission))
}
