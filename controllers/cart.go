package controllers

import (
	"ShopGoApi/common"
	"ShopGoApi/logic"
)

type CartController struct {
	BaseController
	cartLogic logic.CartLogic
}

//@router  /add  [post]
func (c *CartController) PostAdd() {
	cart := new(logic.AddCart)
	cart.UId, _ = c.GetInt("u_id")
	cart.GId, _ = c.GetInt("g_id")
	cart.GNum, _ = c.GetInt("g_num")
	err := c.cartLogic.Add(cart)
	common.HttpResponse(c.Ctx, err)

}
//@router  /all  [get]
func (c *CartController) Get()  {
	u_id ,_:= c.GetInt("u_id")

	data,err := c.cartLogic.GetAll(u_id)

	common.HttpResponseList(c.Ctx,data,err)

}
