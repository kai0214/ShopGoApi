package logic

import (
	"ShopGoApi/models"
)

type (
	GoodsLogic struct {
		goodsModel models.GoodsModel
	}

	AddGoods struct {
		Id       int    `json:"id"       orm:"column(id);auto"`
		Name     string `json:"name"       orm:"column(name)"`
		Describe string `json:"describe"    orm:"column(describe)"`
		Cover    string `json:"cover"       orm:"column(cover)"`
		//CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
		//UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	}

	GoodsDetail struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Describe string `json:"describe"`
		Cover    string `json:"cover"`
	}

	GoodsList []*GoodsItem
	GoodsItem struct {
		Id           int    `json:"id"      `
		Name         string `json:"name"      `
		Describe     string `json:"describe"    `
		Cover        string `json:"cover"  `
		PresentPrice string `json:"present_price"`
		OriginalPrice string `json:"original_price"`


	}
)

//添加
func (l *GoodsLogic) Add(goods *AddGoods) error {
	if err := l.goodsModel.AddGoods(&models.Goods{
		Name:     goods.Name,
		Describe: goods.Describe,
		Cover:    goods.Cover,
		//CreateTime: goods.CreateTime,
		//UpdateTime: goods.CreateTime,
	}); err != nil {
		return err
	}
	return nil
}

//通过ID查找
func (l *GoodsLogic) FindById(id int) (*GoodsDetail, error) {
	m, err := l.goodsModel.FindById(id)
	if err != nil {
		return nil, err
	}
	return &GoodsDetail{
		Id:       m.Id,
		Name:     m.Name,
		Describe: m.Describe,
		Cover:    m.Cover,
	}, nil

}

func (l *GoodsLogic) FindAlOrPage(page int) (*GoodsList, error) {
	date, err := l.goodsModel.FindAlOrPage(page)
	if err != nil {
		return nil, err
	}
	var resp = GoodsList{}
	for _, d := range date {
		resp = append(resp, &GoodsItem{
			Id:       d.Id,
			Name:     d.Name,
			Describe: d.Describe,
			Cover:    d.Cover,
			OriginalPrice:d.OriginalPrice,
			PresentPrice:d.PresentPrice,
		})
	}

	return &resp, nil
}
