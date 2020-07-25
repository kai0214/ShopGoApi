package logic

import (
	"ShopGoApi/common"
	"ShopGoApi/models"
)

type (
	CategoryLogic struct {
		categoryModel models.CategoryModel
	}
	Category struct {
		Id              int             `json:"id" `
		Name            string          `json:"name"  `
		SubCategoryList SubCategoryList `json:"subList"  `
	}

	SubCategory struct {
		Id      int    `json:"id" `
		SubId   int    `json:"sub_id" `
		SubName string `json:"sub_name"  `
	}

	CategoryList    []*Category
	SubCategoryList []*SubCategory
)

func (l *CategoryLogic) AddCategory(name string) error {
	if name == "" {
		return common.NewBaseError(401, "分类名称不能为空")
	}
	if err := l.categoryModel.AddCategory(&models.Category{
		Name: name,
	}); err != nil {
		return err
	}
	return nil
}

func (l *CategoryLogic) AddSubCategory(id int, subName string) error {
	isNil, _ := l.categoryModel.FindCategoryById(id)
	if !isNil {
		return common.NewBaseError(201, "父分类不存在")
	}
	if subName == "" {
		return common.NewBaseError(401, "分类名称不能为空")
	}
	err := l.categoryModel.AddSubCategory(&models.SubCategory{
		Id:      id,
		SubName: subName,
	})
	if err != nil {
		return err
	}
	return nil
}
func (l *CategoryLogic) FindAll() (*CategoryList, error) {
	date, err := l.categoryModel.FindCategoryAll()
	if err != nil {
		return nil, err
	}
	var resp = CategoryList{}
	for _, d := range date {
		var subResp = SubCategoryList{}
		subResp = append(subResp, &SubCategory{
			Id:      d.Id,
			SubId:   0,
			SubName: "全部",
		})
		subData, subErr := l.categoryModel.FindSubCategoryAll(d.Id)
		if subErr != nil {
			return nil, subErr
		}
		for _, sub := range subData {
			subResp = append(subResp, &SubCategory{
				Id:      d.Id,
				SubId:   sub.SubId,
				SubName: sub.SubName,
			})
		}

		resp = append(resp, &Category{
			Id:              d.Id,
			Name:            d.Name,
			SubCategoryList: subResp,
		})

	}

	return &resp, nil
}
