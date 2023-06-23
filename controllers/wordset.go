package controllers

import (
	"encoding/json"

	"github.com/casbin/casibase/object"
	"github.com/casbin/casibase/util"
)

func (c *ApiController) GetGlobalWordsets() {
	c.Data["json"] = object.GetGlobalWordsets()
	c.ServeJSON()
}

func (c *ApiController) GetWordsets() {
	owner := c.Input().Get("owner")

	c.ResponseOk(object.GetWordsets(owner))
}

func (c *ApiController) GetWordset() {
	id := c.Input().Get("id")

	c.ResponseOk(object.GetWordset(id))
}

func (c *ApiController) GetWordsetGraph() {
	id := c.Input().Get("id")
	clusterNumber := util.ParseInt(c.Input().Get("clusterNumber"))
	distanceLimit := util.ParseInt(c.Input().Get("distanceLimit"))

	c.ResponseOk(object.GetWordsetGraph(id, clusterNumber, distanceLimit))
}

func (c *ApiController) GetWordsetMatch() {
	id := c.Input().Get("id")

	c.ResponseOk(object.GetWordsetMatch(id))
}

func (c *ApiController) UpdateWordset() {
	id := c.Input().Get("id")

	var wordset object.Wordset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &wordset)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}
	c.ResponseOk(object.UpdateWordset(id, &wordset))
}

func (c *ApiController) AddWordset() {
	var wordset object.Wordset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &wordset)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.AddWordset(&wordset))
}

func (c *ApiController) DeleteWordset() {
	var wordset object.Wordset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &wordset)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(object.DeleteWordset(&wordset))
}
