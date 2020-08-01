package models

import "github.com/astaxie/beego/orm"

type (
	Cart struct {
		Id   int `json:"id"     orm:"column(id);auto"`
		UId  int `json:"u_id"   orm:"column(u_id)"`
		GId  int `json:"g_id"   orm:"column(g_id)"`
		GNum int `json:"g_num"  orm:"column(g_num)"`
	}

	CartModel struct {
		Cart
	}
)

func init() {
	orm.RegisterModel(new(Cart))
}
func (m *Cart) TabName() string {
	return "`cart`"
}

func (m *CartModel) AddCart(cart *Cart) error {
	_, err := orm.NewOrm().Insert(cart)
	if err != nil {
		return err
	}
	return nil
}

func (m *CartModel) UpdateCart(cart *Cart) error {
	_, err := orm.NewOrm().Update(cart)
	if err != nil {
		return err
	}
	return nil
}

/**
判断是否存在当前商品
*/
func (m *CartModel) GoodsFindId(id int) (bool, int, error) {
	c := new(Goods)
	if err := orm.NewOrm().Raw("SELECT * FROM "+"`goods`"+" WHERE id = ? ", id).QueryRow(c);
		err != nil {
		return false, 0, err
	}
	if c.Id != 0 {
		return true, c.Num, nil
	}
	return false, 0, nil
}

/**
判断是否存在当前账号
*/
func (m *CartModel) UserFindById(id int) (bool, error) {
	c := new(Account)
	if err := orm.NewOrm().Raw("SELECT * FROM "+"account"+" WHERE id = ? ", id).QueryRow(c);
		err != nil {
		return false, err
	}
	if c.Phone != "" {
		return true, nil
	}
	return false, nil
}

/**
通过用户和商品ID  获取购物车id
*/
func (m *CartModel) FindUserAndGoods(u_id, g_id int) (*Cart, error) {
	c := new(Cart)
	if err := orm.NewOrm().Raw("SELECT * FROM "+m.TabName()+" WHERE u_id = ? and g_id =? ", u_id, g_id).QueryRow(c);
		err != nil {
		return nil, err
	}
	return c, nil
}

/**
获取购物车列表
*/
func (m *CartModel) FindAll(u_id int) ([]*Cart, error) {
	cs := make([]*Cart, 0)
	if _, err := orm.NewOrm().Raw("SELECT * FROM "+m.TabName()+" WHERE u_id = ? ", u_id).QueryRows(&cs);
		err != nil {
		return nil, err
	}
	return cs, nil
}
