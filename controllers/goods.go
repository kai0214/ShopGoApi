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

//@router /addGoods [post]
func (c *GoodsController) AddGoods() {
	goods := new(logic.AddGoods)
	goods.Name = c.GetString("name")
	goods.Describe = c.GetString("describe")
	goods.Cover = c.GetString("cover")
	valid := new(validation.Validation)
	valid.Required(goods.Name, "name").Message("商品名称不能为空")
	valid.Required(goods.Describe, "describe").Message("商品描述不能为空")
	valid.Required(goods.Cover, "cover").Message("商品封面不能为空")
	if valid.HasErrors() {
		for _, e := range valid.Errors {
			common.HttpResponse(c.Ctx,
				common.NewBaseError(common.ErrFromValid, e.Message))
			return
		}
	}

	err := c.goodsLogic.Add(goods)
	common.HttpResponse(c.Ctx, err)
}

//@router /detail  [get]
func (c *GoodsController) GetDetail() {
	id, _ := c.GetInt("id")
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
	data, err := c.goodsLogic.FindById(id)
	common.HttpResponseData(c.Ctx, data, err)


}

//@router  /list  [get]
func (c *GoodsController) GetPageList() {
	page, _ := c.GetInt("page")
	data, err := c.goodsLogic.FindByPage(page)
	common.HttpResponseList(c.Ctx, data, err)
}
