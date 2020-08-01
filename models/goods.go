package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type (
	Goods struct {
		Id       int    `json:"id"       orm:"column(id);auto"`
		Name     string `json:"name"       orm:"column(g_name)"`
		Describe string `json:"describe"    orm:"column(g_describe)"`
		Cover    string `json:"cover"       orm:"column(cover)"`
		Num      int    `json:"goodsNum" orm:"column(num)"`
		PresentPrice  float64 `json:"present_price" orm:"column(present_price)"`
		OriginalPrice float64 `json:"original_price" orm:"column(original_price)"`
		Category      int     `json:"category" orm:"column(category)"`
		SubCategory   int     `json:"subCategory" orm:"column(sub_category)"`
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

func (m *GoodsModel) FindAllOrPage(page int) ([]*Goods, error) {
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

/**
通过子分类获取分页商品
*/
func (m *GoodsModel) FindSubCategoryPage(category, subCategory, page int) ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().Raw("SELECT * FROM "+m.TableName()+" WHERE category = ? and sub_category =?   ORDER BY create_time DESC LIMIT ?,?", category, subCategory, page*10, (page+1)*10).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}

/**
通过父分类获取分页商品
*/

func (m *GoodsModel) FindCategoryPage(category, page int) ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().Raw("SELECT * FROM "+m.TableName()+" WHERE category = ?   ORDER BY create_time DESC LIMIT ?,?", category, page*10, (page+1)*10).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}

/**
通过子分类获取全部商品
*/
func (m *GoodsModel) FindSubCategoryAll(category, subCategory int) ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM "+m.TableName()+"WHERE category = ? and sub_category =?   ORDER BY create_time DESC", category, subCategory).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}

/**
通过父分类获取全部商品
*/
func (m *GoodsModel) FindCategoryAll(category int) ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM "+m.TableName()+"WHERE category = ?    ORDER BY create_time DESC", category).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}

func (m *GoodsModel) FindCategoryAllOrPage(category, subCategory, page int) ([]*Goods, error) {
	if page > 0 {
		if subCategory == 0 {
			return m.FindCategoryPage(category, page-1)
		} else {
			return m.FindSubCategoryPage(category, subCategory, page-1)
		}

	} else {
		if subCategory == 0 {
			return m.FindCategoryAll(category)
		} else {
			return m.FindSubCategoryAll(category, subCategory)
		}
	}
}

/**
判断是否存在当前父分类
*/
func (m *GoodsModel) CategoryFindId(category int) (bool, error) {
	c := new(Category)
	if err := orm.NewOrm().Raw("SELECT * FROM "+"`category`"+" WHERE id = ? ", category).QueryRow(c);
		err != nil {
		return false, err
	}
	if c.Id != 0 {
		return true, nil
	}
	return false, nil
}

/**
判断是否存在当前父分类
*/
func (m *GoodsModel) SubCategoryFindId(category, subCategory int) (bool, error) {
	c := new(SubCategory)
	if err := orm.NewOrm().Raw("SELECT * FROM "+"`sub_category`"+" WHERE id = ? and  sub_id = ?", category, subCategory).QueryRow(c);
		err != nil {
		return false, err
	}
	if c.SubId != 0 {
		return true, nil
	}
	return false, nil
}
