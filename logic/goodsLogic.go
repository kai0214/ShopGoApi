package logic

import "ShopGoApi/models"

type (
	GoodsLogic struct {
		goodsModel models.GoodsModel
	}

	AddGoods struct {
		Id       int    `json:"id"      `
		Name     string `json:"name"      `
		Describe string `json:"describe"    `
		Cover    string `json:"cover"  `
	}

	GoodsDetail struct {
		Id       int    `json:"id"      `
		Name     string `json:"name"      `
		Describe string `json:"describe"    `
		Cover    string `json:"cover"  `
	}

	GoodsList []*GoodsItem
	GoodsItem struct {
		Id       int    `json:"id"      `
		Name     string `json:"name"      `
		Describe string `json:"describe"    `
		Cover    string `json:"cover"  `
	}
)

//添加
func (l *GoodsLogic) Add(goods *AddGoods) error {
	if err := l.goodsModel.AddGoods(&models.Goods{
		Name:     goods.Name,
		Describe: goods.Describe,
		Cover:    goods.Cover,
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

func (l *GoodsLogic) FindByPage(page int) (*GoodsList, error) {
	date, err := l.goodsModel.FindByPage(page)
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
		})
	}
	return &resp, nil
}
