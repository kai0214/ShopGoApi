package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type (
	Account struct {
		Id         int       `json:"id"            orm:"column(id);auto"`
		Phone      string    `json:"phone"         orm:"column(phone)"`
		Password   string    `json:"password"           orm:"column(pwd)"`
		CreateTime time.Time `json:"create_time"   orm:"column(create_time)"`
		UpdateTime time.Time `json:"update_time"   orm:"column:(update_time)"`
	}

	AccountModel struct {
		Account
	}
)

func init() {
	orm.RegisterModel(new(Account))
}

func (m *Account) TableName() string {
	return "`account`"

}

/**
添加账号
*/
func (m *AccountModel) AddAccount(a *Account) error {
	//if _, err := orm.NewOrm().Insert(a); err != nil {
	//	return err
	//}
	_, err := orm.NewOrm().Raw("insert into "+m.TableName()+" (phone,password) values(?,?)", a.Phone, a.Password).Exec()
	if err != nil {
		return err
	}
	return nil
}

/**
判断是否存在当前账号
*/
func (m *AccountModel) FindByPhone(phone string) (bool, error) {
	c := new(Account)
	if err := orm.NewOrm().Raw("SELECT * FROM "+m.TableName()+" WHERE phone = ? ", phone).QueryRow(c);
		err != nil {
		return false, err
	}
	if c.Phone != "" {
		return true, nil
	}
	return false, nil
}

/**
通过账号和密码查询   用户登陆
*/
func (m *AccountModel) FindPhoneAndPassword(phone, password string) (*Account, error) {
	c := new(Account)
	if err := orm.NewOrm().Raw("SELECT * FROM "+m.TableName()+" WHERE phone = ? and password =? ", phone, password).QueryRow(c);
		err != nil {
		return c, err
	}
	return c, nil
}
