package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	_"github.com/go-sql-driver/mysql"

)

type UserInfo struct {
	Id int
	Name string
	Phone int
	OrderInfo  *OrderInfo  `orm:"rel(fk)"`    //设置一对多关系


}

type OrderInfo struct {
	Id int
	Gsname string
	Gsprice string
	Gstime time.Time
	UserInfo        []*UserInfo `orm:"reverse(many)"` // 设置一对多的反向关系
}

//type User struct {
//	Id        int64   `json:"id" `
//	Name      string  `json:"name,omitempty" orm:"size(50)"`
//	Passwords string  `json:"passwords" orm:"size(32)"`
//	Baby      []*Baby `json:"baby" orm:"reverse(many)"`
//}
//type Baby struct {
//	Id int64
//	Name string `json:"name" orm:"size(50)"`
//	User *User `json:"user" orm:"rel(fk);index"`
//}
/*
 会员（用户） -> 文章：一对多
文章 -> 文章分类：多对一
文章 -> 评论：一对多
说明：beego的orm使用时，外键id在关联查询时会默认添加一个"_id"结尾，
比如：文章表对应的作者id，orm在关联查询时会默认查询xxx_id，其中xxx为struct中定义的json字段全称，
这样的话最好定义外键id时直接写成xxx_id形式，然后struct的字段的json tag写成xxx即可。

*/

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Mobile    string     `json:"mobile"`
	Age       int        `json:"age"`
	Sex       bool       `json:"sex"`
	Articles  []*Article `orm:"reverse(many)"` //反向一对多
}

type Article struct {
	Id       int          `json:"id"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Addtime  int          `json:"addtime"`
	Type     *Articletype `json:"type" orm:"rel(fk)"`
	User     *User        `json:"user" orm:"rel(fk)"` //一对多
	//Type     *Articletype `json:"type" orm:"rel(fk)"`
	//Comments []*Comment   `orm:"reverse(many)"` //反向一对多关联
}

type Articletype struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Orderno  int        `json:"orderno"`
	Articles []*Article `orm:"reverse(many)"`
}

type Hello struct {
	// TODO:问题10：这样写实际生成的表是什么样子
	CommonModel
	// TODO:问题11：orm:"xxxx" 与 json 的用途分别是什么
	//orm:"xxxx" 对属性进行标注，用于解析  json 改变json格式的key值为name

	Name   string `orm:"column(user_name)" json:"name"`
	Age    int
	Gender int
}

type CommonModel struct {
	Id         int64
	CreateTime int64
	UpdateTime int64
	Deleted    bool
}








//type User struct {
//	Id int
//	UserName string
//	Passwd string
//	Articles []*article`orm:"rel(m2m)"`  //用户和文章多对多 一片文章可以被多个用户看一个用户可以看多篇文章
//}

//文章表和文章类型表示一对多
//type Article struct {
//	Id2 int `orm:"pk;auto"`
//	Title string `orm:"size(20)"`
//	Content string `orm:"size(500)"`
//	Img string `orm:"size(50);null"`
//	Count int `orm:"default(0)"`
//	ArticleType *ArticleType `orm:"rel(fk)"`

//Users []*User `orm:"reverse(many)"`
//}
//
//type ArticleType struct {
//	Id int
//	TypeName string `orm:"size(20)"`
//	Article[]*Article `orm:"reverse(many)"`
//}







func init(){
	orm.RegisterDataBase("default","mysql","root:1234@tcp(127.0.0.1:3306)/test02?charset=utf8")
	orm.RegisterModel(new(UserInfo),new(OrderInfo),new(User),new(Article),new(Articletype),new(Hello))
	orm.RunSyncdb("default",false,true)
}