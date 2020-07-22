package logic

import (
	"ShopGoApi/common"
	"ShopGoApi/models"
	"time"
)

type (
	AccountLogic struct {
		accountModel models.AccountModel
	}

	AddAccount struct {
		Id         int       `json:"id"`
		Phone      string    `json:"phone"`
		Password   string    `json:"password"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
	}

	Info struct {
		Id       int    `json:"id"`
		Phone    string `json:"phone"`
		Password string `json:"password,omitempty"`
	}
)

/**
添加账号
*/
func (l *AccountLogic) Add(a *AddAccount) error {
	if a.Phone == "" {
		return common.NewBaseError(401, "账号不能为空")
	}
	if a.Password == "" {
		return common.NewBaseError(401, "密码不能为空")
	}
	isNil, _ := l.accountModel.FindByPhone(a.Phone)
	if isNil {
		return common.NewBaseError(201, "账号已存在")
	}
	if err := l.accountModel.AddAccount(&models.Account{
		Phone:    a.Phone,
		Password: a.Password,
	}); err != nil {
		return err
	}
	return nil
}

/**
用户登陆
*/
func (l *AccountLogic) FindPhoneAndPassword(phone, password string) (*Info, error) {
	if phone == "" {
		return nil, common.NewBaseError(401, "账号不能为空")
	}
	if password == "" {
		return nil, common.NewBaseError(401, "密码不能为空")
	}
	data, err := l.accountModel.FindPhoneAndPassword(phone, password)
	if err != nil {
		return nil, common.NewBaseError(400, "账号或密码错误")
	}
	return &Info{
		Id:    data.Id,
		Phone: data.Phone,
	}, nil
}
