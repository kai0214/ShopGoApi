package controllers

import (
	"ShopGoApi/common"
	"ShopGoApi/logic"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type GoodsController struct {
	BaseController
	goodsLogic logic.GoodsLogic
}

//@router /add [post]
func (c *GoodsController) AddGoods() {
	goods := new(logic.AddGoods)
	goods.Name = c.GetString("name")
	goods.Describe = c.GetString("describe")
	goods.Cover = c.GetString("cover")
	goods.Category, _ = c.GetInt("category")
	goods.Num, _ = c.GetInt("num")
	goods.SubCategory, _ = c.GetInt("sub_category")
	goods.PresentPrice, _ = c.GetFloat("present_price")
	goods.OriginalPrice, _ = c.GetFloat("original_price")

	err := c.goodsLogic.Add(goods)
	common.HttpResponse(c.Ctx, err)
}

//@router /detail  [get]
func (c *GoodsController) GetDetail() {
	id, _ := c.GetInt("id")
	u_id, _ := c.GetInt("u_id")
	fmt.Print("----Id:", id)
	valid := new(validation.Validation)
	valid.Required(id, "id").Message("Id异常")
	if valid.HasErrors() {
		for _, e := range valid.Errors {
			common.HttpResponse(c.Ctx,
				common.NewBaseError(common.ErrFromValid, e.Message))
			return
		}
	}
	data, err := c.goodsLogic.FindById(u_id, id)
	fmt.Print(data)
	common.HttpResponseData(c.Ctx, data, err)

}

//@router  /list  [get]
func (c *GoodsController) GetPageList() {
	page, _ := c.GetInt("page")
	data, err := c.goodsLogic.FindAlOrPage(page)
	common.HttpResponseList(c.Ctx, data, err)
}

//@router  /categoryGoodsList  [get]
func (c *GoodsController) GetCategoryGoodsList() {
	page, _ := c.GetInt("page")
	category, _ := c.GetInt("category")
	subCategory, _ := c.GetInt("sub_category")
	data, err := c.goodsLogic.FindCategoryAlOrPage(category, subCategory, page)
	common.HttpResponseList(c.Ctx, data, err)
}
