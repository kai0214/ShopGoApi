package controllers

import (
	"ShopGoApi/common"
	"ShopGoApi/logic"
)

type (
	CategoryController struct {
		BaseController
		categoryLogic logic.CategoryLogic
	}
)

//@router /list   [get]
func (c *CategoryController) Get() {
	data, err := c.categoryLogic.FindAll()
	common.HttpResponseData(c.Ctx, data, err)

}

//@router /add [post]
func (c *CategoryController) Post() {
	name := c.GetString("name")
	err := c.categoryLogic.AddCategory(name)
	common.HttpResponse(c.Ctx, err)

}

//@router /addSub [post]
func (c *CategoryController) PostSub() {
	name := c.GetString("name")
	id, _ := c.GetInt("id")
	err := c.categoryLogic.AddSubCategory(id, name)
	common.HttpResponse(c.Ctx, err)
}
