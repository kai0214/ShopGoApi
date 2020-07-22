# Beego开发Api 

## bee 创建连接数据库的API
bee api appname  -conn="root:q123456@tcp(localhost:3306)/test"
### 创建mod
 go mod init
### 添加已有的库
go build ./...
### 使用API自动文档
bee run -gendoc=true -downdoc=true

访问 localhost:8080/swagger

### Json使用
 >1.`json:"id ,-"`  忽略字段 

 >2.`json:"status,optional"`  可选字段

 >3.`json:"input,omitempty"`  当字段为空时忽略此字段

### orm 操作数据库
1. 部分更新数据 
 
  orm.NewOrm().Update(a, "pwd", "update_time")
### router 注解
```
// @Title Create   标题
// @Description create object   说明
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
```
