package controllers

import (
	"ShopGoApi/common"
	"ShopGoApi/logic"
)

type AccountController struct {
	BaseController
	accountLogic logic.AccountLogic
}

//@router   /register    [post]
func (c *AccountController) Register() {
	account := new(logic.AddAccount)
	account.Phone = c.GetString("phone")
	account.Password = c.GetString("password")
	err := c.accountLogic.Add(account)
	common.HttpResponse(c.Ctx, err)
}

//@router   /login   [post]
func (c *AccountController) Login() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	info, err := c.accountLogic.FindPhoneAndPassword(phone, password)
	common.HttpResponseData(c.Ctx, info, err)
}
