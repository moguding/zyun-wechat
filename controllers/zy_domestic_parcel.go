package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"zyun-wechat-api/models"

	"github.com/astaxie/beego"
)

// oprations for ZyDomesticParcel
type ZyDomesticParcelController struct {
	beego.Controller
}

func (c *ZyDomesticParcelController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create ZyDomesticParcel
// @Param	body		body 	models.ZyDomesticParcel	true		"body for ZyDomesticParcel content"
// @Success 200 {int} models.ZyDomesticParcel.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ZyDomesticParcelController) Post() {
	var v models.ZyDomesticParcel
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if id, err := models.AddZyDomesticParcel(&v); err == nil {
		c.Data["json"] = map[string]int64{"id": id}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Get
// @Description get ZyDomesticParcel by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ZyDomesticParcel
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ZyDomesticParcelController) GetOne() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetZyDomesticParcelById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJson()
}

// @Title Get All
// @Description get ZyDomesticParcel
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ZyDomesticParcel
// @Failure 403
// @router / [get]
func (c *ZyDomesticParcelController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllZyDomesticParcel(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJson()
}

// @Title Update
// @Description update the ZyDomesticParcel
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ZyDomesticParcel	true		"body for ZyDomesticParcel content"
// @Success 200 {object} models.ZyDomesticParcel
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ZyDomesticParcelController) Put() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.ZyDomesticParcel{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateZyDomesticParcelById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}

// @Title Delete
// @Description delete the ZyDomesticParcel
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ZyDomesticParcelController) Delete() {
	idStr := c.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteZyDomesticParcel(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJson()
}
