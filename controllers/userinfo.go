package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"onetomany/models"

)

type SearchController struct {
	beego.Controller
}

//func (this *SearchController)Get()  {
//	o := orm.NewOrm()
//	var user models.OrderInfo
//	qs := o.QueryTable(user)
//	qs.Filter("id",1)
//	//qs.Filter("profile__age",18).RelatedSel().All(&user)
//	beego.Info(user)
//	beego.Info("hello")
//}

func(this *SearchController)Get(){

	//根据用户查询所有文章

	//var articles []*models.Article
	//orm.NewOrm().QueryTable("article").Filter("user", 1).RelatedSel().All(&articles)
	//for _, v := range articles {
	//	fmt.Println(v.Title)
	//}

	//var articles []*models.Article
	//orm := orm.NewOrm()
	//n,err := orm.QueryTable("article").Filter("user__name","唐家三少").All(&articles)
	//
	//if err != nil{
	//	this.Ctx.WriteString("查询失败")
	//	return
	//}
	//fmt.Println("n=",n)
	//for _,article := range articles{
	//	fmt.Println("article=",article)
	//}
	//this.Ctx.WriteString("查询成功")


//2、根据文章查询对应用户
//	var user models.User
//	err := orm.NewOrm().QueryTable("user").Filter("Name", "金庸").Limit(1).One(&user)
//	if err == nil {
//		fmt.Println(user)
//	}

	var user models.User
	orm := orm.NewOrm()
	err := orm.QueryTable("user").Filter("articles__title","天龙八部").One(&user)

	if err != nil{
		this.Ctx.WriteString("查询出错")
		return
	}
	fmt.Println("user=",user)
	this.Ctx.WriteString("查询成功")

}

//func (this *SearchController) UpdateUser() {
//	user := models.Hello{Name: "xxx", Age: 1, CommonModel: models.CommonModel{CreateTime: 2, UpdateTime: 3}}
//
//	updateName(&user)
//
//	updateTime(user)
//
//	// TODO:问题12：这里的实际结果是什么
//	//this.Success(user)
//}
//
//func updateName(user *models.Hello) {
//	user.Name = "update" + user.Name
//}
//
//func updateTime(user models.Hello) {
//	user.UpdateTime = time.Now().Unix()
//}


//type BaseController struct {
//	beego.Controller
//}
//
//
//func (this *BaseController) Success(data interface{}) {
//	this.Data["json"] = beego.M{"code": 200, "data": data}
//	// TODO:问题2：这里可以传参，用途是什么
//	this.ServeJSON()
//	this.StopRun()
//}

func (this *UserController) UpdateUser() {
	user := models.User{Name: "xxx", Age: 1, CommonModel: models.CommonModel{CreateTime: 2, UpdateTime: 3}}

	updateName(&user)

	updateTime(user)

	// TODO:问题12：这里的实际结果是什么
	//JSON序列化输出user
	this.Success(user)
}