package logic

import (
	"ShopGoApi/common"
	"ShopGoApi/models"
)

type (
	GoodsLogic struct {
		goodsModel models.GoodsModel
	}

	AddGoods struct {
		Id            int     `json:"id"       orm:"column(id);auto"`
		Name          string  `json:"name"       orm:"column(name)"`
		Describe      string  `json:"describe"    orm:"column(describe)"`
		Cover         string  `json:"cover"       orm:"column(cover)"`
		Category      int     `json:"category" orm:"column(category)"`
		SubCategory   int     `json:"subCategory" orm:"column(sub_category)"`
		PresentPrice  float64 `json:"present_price" orm:"column(present_price)"`
		OriginalPrice float64 `json:"original_price" orm:"column(original_price)"`
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
		Id            int     `json:"id"      `
		Name          string  `json:"name"      `
		Describe      string  `json:"describe"    `
		Cover         string  `json:"cover"  `
		PresentPrice  float64 `json:"present_price"`
		OriginalPrice float64 `json:"original_price"`
	}
)

//添加
func (l *GoodsLogic) Add(goods *AddGoods) error {
	if goods.Name == "" {
		return common.NewBaseError(201, "商品名称不能为空")
	}
	if goods.Cover == "" {
		return common.NewBaseError(201, "商品封面不能为空")
	}
	if goods.Describe == "" {
		return common.NewBaseError(201, "商品描述不能为空")
	}
	if goods.OriginalPrice == 0 {
		return common.NewBaseError(201, "商品价格不能为空")
	}
	isNil, _ := l.goodsModel.CategoryFindId(goods.Category)
	if !isNil {
		return common.NewBaseError(201, "商品父分类不存在")
	}
	isSubNil, _ := l.goodsModel.SubCategoryFindId(goods.Category, goods.SubCategory)
	if !isSubNil {
		return common.NewBaseError(201, "商品子分类不存在")
	}
	if err := l.goodsModel.AddGoods(&models.Goods{
		Name:          goods.Name,
		Describe:      goods.Describe,
		Cover:         goods.Cover,
		Category:      goods.Category,
		SubCategory:   goods.SubCategory,
		OriginalPrice: goods.OriginalPrice,
		PresentPrice:  goods.PresentPrice,
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
	date, err := l.goodsModel.FindAllOrPage(page)
	if err != nil {
		return nil, err
	}
	var resp = GoodsList{}
	for _, d := range date {
		resp = append(resp, &GoodsItem{
			Id:            d.Id,
			Name:          d.Name,
			Describe:      d.Describe,
			Cover:         d.Cover,
			OriginalPrice: d.OriginalPrice,
			PresentPrice:  d.PresentPrice,
		})
	}

	return &resp, nil
}

func (l *GoodsLogic) FindCategoryAlOrPage(category, subCategory, page int) (*GoodsList, error) {
	date, err := l.goodsModel.FindCategoryAllOrPage(category, subCategory, page)
	if err != nil {
		return nil, err
	}
	var resp = GoodsList{}
	for _, d := range date {
		resp = append(resp, &GoodsItem{
			Id:            d.Id,
			Name:          d.Name,
			Describe:      d.Describe,
			Cover:         d.Cover,
			OriginalPrice: d.OriginalPrice,
			PresentPrice:  d.PresentPrice,
		})
	}

	return &resp, nil
}
