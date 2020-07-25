package models

import "github.com/astaxie/beego/orm"

type (
	Category struct {
		Id   int    `json:"id" orm:"column(id);auto"`
		Name string `json:"name"  orm:"column(c_name)"`
	}

	SubCategory struct {
		Id      int    `json:"id" orm:"column(id)"`
		SubId   int    `json:"sub_id" orm:"column(sub_id);auto"`
		SubName string `json:"sub_name" orm:"column(sub_name)" `
	}

	CategoryModel struct {
		Category
		SubCategory
	}
)

func (c *Category) TableName() string {
	return "`category`"
}

func (c *SubCategory) SubTableName() string {
	return "`sub_category`"
}

//添加分类
func (m *CategoryModel) AddCategory(c *Category) error {
	_, err := orm.NewOrm().Raw("insert into "+m.TableName()+" (c_name) values(?)", c.Name).Exec()
	if err != nil {
		return err
	}
	return nil
}

//查询分类是否存在
func (m *CategoryModel) FindCategoryById(id int) (bool, error) {
	c := new(Category)
	if err := orm.NewOrm().Raw("SELECT * FROM "+m.TableName()+" WHERE id = ? ", id).QueryRow(c);
		err != nil {
		return false, err
	}
	if c.Id != 0 {
		return true, nil
	}
	return false, nil
}

//添加子类
func (m *CategoryModel) AddSubCategory(sub *SubCategory) error {
	_, err := orm.NewOrm().Raw("insert into "+m.SubTableName()+" (sub_name,id) values(?,?)", sub.SubName, sub.Id).Exec()
	if err != nil {
		return err
	}
	return nil
}

//获取全部分类
func (m *CategoryModel) FindCategoryAll() ([]*Category, error) {
	cs := make([]*Category, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM " + m.TableName() + "").
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}

//获取全部子分类
func (m *CategoryModel) FindSubCategoryAll(id int) ([]*SubCategory, error) {
	cs := make([]*SubCategory, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM "+m.SubTableName()+" WHERE id = ? ", id).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}
