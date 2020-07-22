package models
import "github.com/astaxie/beego/orm"

type (
	Goods struct {
		Id       int    `json:"id"       orm:"column(id);auto"`
		Name     string `json:"name"       orm:"column(name)"`
		Describe string `json:"describe"    orm:"column(describe)"`
		Cover    string `json:"cover"       orm:"column(cover)"`
	}

	GoodsModel struct {
		Goods
	}
)

func init() {
	orm.RegisterModel(new(Goods))
}

func (g *Goods) TableName() string {
	return "`goods`"
}

/**
添加商品
*/
func (m *GoodsModel) AddGoods(g *Goods) error {
	if _, err := orm.NewOrm().Insert(g); err != nil {
		return err
	}
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

/**
分页获取商品
*/
func (m *GoodsModel) FindByPage(page int) ([]*Goods, error) {
	cs := make([]*Goods, 0)
	if _, err := orm.NewOrm().
		Raw("SELECT * FROM "+m.TableName()+"  ORDER BY create_time DESC LIMIT ?,2", page-1).
		QueryRows(&cs);
		nil != err {
		return nil, err
	}
	return cs, nil
}
