package logic

import (
	"ShopGoApi/common"
	"ShopGoApi/models"
)

type (
	CartLogic struct {
		cartModel  models.CartModel
		goodsModel models.GoodsModel
	}
	AddCart struct {
		UId  int `json:"u_id"   orm:"column(u_id)"`
		GId  int `json:"g_id"   orm:"column(g_id)"`
		GNum int `json:"g_num"  orm:"column(g_num)"`
	}

	CartGoodsList []*CartGoods
	CartGoods     struct {
		Id     int    `json:"id"`
		GId    int    `json:"g_id"`
		GName  string `json:"g_name"`
		GNum   int    `json:"g_num"`
		GCover string `json:"g_cover"`
	}
)

func (l *CartLogic) Add(cart *AddCart) error {
	userIsNil, _ := l.cartModel.UserFindById(cart.UId)
	if !userIsNil {
		return common.NewBaseError(401, "登陆失效")
	}
	isNil, num, _ := l.cartModel.GoodsFindId(cart.GId)
	if !isNil {
		return common.NewBaseError(201, "商品不存在")
	}
	data, _ := l.cartModel.FindUserAndGoods(cart.UId, cart.GId)

	if cart.GNum <= 0 && data != nil {
		cart.GNum = data.GNum + 1
	} else {
		cart.GNum = 1
	}
	if cart.GNum > num {
		return common.NewBaseError(201, "商品数量超出库存")
	}
	if data != nil {
		err := l.cartModel.UpdateCart(&models.Cart{
			Id:   data.Id,
			UId:  cart.UId,
			GId:  cart.GId,
			GNum: cart.GNum,
		})
		if err != nil {
			return err
		}
		return nil
	} else {
		err := l.cartModel.AddCart(&models.Cart{
			UId:  cart.UId,
			GId:  cart.GId,
			GNum: cart.GNum,
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func (l *CartLogic) GetAll(u_id int) (*CartGoodsList, error) {
	userIsNil, _ := l.cartModel.UserFindById(u_id)
	if !userIsNil {
		return nil, common.NewBaseError(401, "登陆失效")
	}
	cartData, err := l.cartModel.FindAll(u_id)
	if err != nil {
		return nil, err
	}
	var resp = CartGoodsList{}
	for _, d := range cartData {
		goods, err := l.goodsModel.FindById(d.GId)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &CartGoods{
			Id:     d.Id,
			GId:    d.GId,
			GName:  goods.Name,
			GNum:   d.GNum,
			GCover: goods.Cover,
		})
	}

	return &resp, nil

}
