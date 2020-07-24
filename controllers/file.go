package controllers

import (
	"ShopGoApi/common"
	"ShopGoApi/logic"
	"math/rand"
	"os"
	"path"
	"time"
)

type FileController struct {
	BaseController
	fileLogic logic.FileLogic
}

//@router /    [post]
func (c *FileController) Post() {

	file, fileHeader, fileErr := c.GetFile("file") //获取上传的文件
	if fileErr != nil {
		common.HttpResponse(c.Ctx, common.NewBaseError(201, "文件异常"))
		return
	}
	ext := path.Ext(fileHeader.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		common.HttpResponse(c.Ctx, common.NewBaseError(201, "后缀名不符合上传要求"))
		return
	}
	//创建目录
	uploadDir := "static/"
	err := os.MkdirAll(uploadDir, 0777)
	if err != nil {

	}
	fileName := GetRandomString(20) + ext //将文件信息头的信息赋值给filename变量
	fpath := uploadDir + fileName
	defer file.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", fpath)
	if err != nil {
		common.HttpResponse(c.Ctx, err)
		return
	}
	common.HttpResponseData(c.Ctx, logic.File{FilePath: fpath}, err)

}

func GetRandomString(l int) string {
	str := "0123456789qwertyuiopasdfghjklzxcvbnm"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
