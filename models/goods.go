package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type (
	Goods struct {
		Id            int    `json:"id"       orm:"column(id);auto"`
		Name          string `json:"name"       orm:"column(g_name)"`
		Describe      string `json:"describe"    orm:"column(g_describe)"`
		Cover         string `json:"cover"       orm:"column(cover)"`
		PresentPrice  string `json:"present_price" orm:"column(present_price)"`
		OriginalPrice string `json:"original_price" orm:"column(original_price)"`
		//CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
		//UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	}

	GoodsModel struct {
		Goods
	}
)

func init() {
	orm.RegisterModel(new(Goods))
}

func (g *Goods) TableName() string {
	return "goods"
}

/**
添加商品
*/
func (m *GoodsModel) AddGoods(g *Goods) error {
	if _, err := orm.NewOrm().Insert(g); err != nil {
		fmt.Printf(err.Error())
		return err
	}

	//fmt.Println("---tableName  ", m.TableName(), g.Name, g.Describe, g.Cover)
	//_, err := orm.NewOrm().Raw("INSERT INTO "+m.TableName()+"(name,describe,cover) VALUES (?,?,?) ", g.Name, g.Describe, g.Cover).Exec()
	//if err != nil {
	//	return err
	//}
	return nil
}

/**
更新
*/
func (m *GoodsModel) UpdateById(g *Goods) error {
	if _, err := orm.NewOrm().Update(g); err != nil {
		return err
	}
	return nil
}

/**
通过id 获取商品
*/
func (m *GoodsModel) FindById(id int) (*Goods, error) {
	c := new(Goods)
	if err := orm.NewOrm().Raw("SELECT * FROM "+m.TableName()+" WHERE id = ? ", id).QueryRow(c);
		err != nil {
		return nil, err
	}
	return c, nil
}

func (m *GoodsModel) FindAlOrPage(page int) ([]*Goods, error) {
	if page > 0 {
		return m.FindByPage(page - 1)
	}
	return m.FindAll()
}

/**
分页获取商品
*/
func (m *GoodsModel) FindByPage(page int) ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM "+m.TableName()+"  ORDER BY create_time DESC LIMIT ?,?", page*10, (page+1)*10).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}

/**
获取全部
*/

func (m *GoodsModel) FindAll() ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM " + m.TableName() + "  ORDER BY create_time DESC").
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}
