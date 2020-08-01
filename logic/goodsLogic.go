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
		Num           int     `json:"num"`
		Category      int     `json:"category" orm:"column(category)"`
		SubCategory   int     `json:"subCategory" orm:"column(sub_category)"`
		PresentPrice  float64 `json:"present_price" orm:"column(present_price)"`
		OriginalPrice float64 `json:"original_price" orm:"column(original_price)"`
		//CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
		//UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	}

	GoodsDetail struct {
		Id            int      `json:"id"`
		Name          string   `json:"name"`
		Describe      string   `json:"describe"`
		Cover         string   `json:"cover"`
		Category      int      `json:"category" orm:"column(category)"`
		SubCategory   int      `json:"subCategory" orm:"column(sub_category)"`
		PresentPrice  float64  `json:"present_price" orm:"column(present_price)"`
		OriginalPrice float64  `json:"original_price" orm:"column(original_price)"`
		Images        []string `json:"images"`
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
	if goods.Num <= 0 {
		return common.NewBaseError(201, "商品数量异常")
	}
	if err := l.goodsModel.AddGoods(&models.Goods{
		Name:          goods.Name,
		Describe:      goods.Describe,
		Cover:         goods.Cover,
		Category:      goods.Category,
		SubCategory:   goods.SubCategory,
		Num:           goods.Num,
		OriginalPrice: common.Decimal(goods.OriginalPrice),
		PresentPrice:  common.Decimal(goods.PresentPrice),

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
	imgs := []string{"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596108392729&di=d54b9acd996780253028ebba3ed64d9d&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2F201604%2F23%2F002205xqdkj84gnw4oi85v.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596108419697&di=5e3091ccb832c3a3951009558aac64a2&imgtype=0&src=http%3A%2F%2Fimg.pconline.com.cn%2Fimages%2Fupload%2Fupc%2Ftx%2Fwallpaper%2F1305%2F03%2Fc0%2F20496484_1367549048818.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596108419697&di=6aae680f0959c982a1975a655493324d&imgtype=0&src=http%3A%2F%2Fattach.bbs.miui.com%2Fforum%2F201408%2F28%2F173210raohlf66afhfaffj.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596108419689&di=3c1a7363078e4e9104bd2ef03bffca8b&imgtype=0&src=http%3A%2F%2Fimg.pconline.com.cn%2Fimages%2Fupload%2Fupc%2Ftx%2Fwallpaper%2F1309%2F05%2Fc4%2F25279801_1378348357336.jpg"}

	if err != nil {
		return nil, err
	}
	return &GoodsDetail{
		Id:            m.Id,
		Name:          m.Name,
		Describe:      m.Describe,
		Cover:         m.Cover,
		OriginalPrice: common.Decimal(m.OriginalPrice),
		PresentPrice:  common.Decimal(m.PresentPrice),
		Category:      m.Category,
		SubCategory:   m.SubCategory,
		Images:        imgs,
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
			OriginalPrice: common.Decimal(d.OriginalPrice),
			PresentPrice:  common.Decimal(d.PresentPrice),
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
			OriginalPrice: common.Decimal(d.OriginalPrice),
			PresentPrice:  common.Decimal(d.PresentPrice),
		})
	}

	return &resp, nil
}
