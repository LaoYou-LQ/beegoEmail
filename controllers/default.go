package controllers

import (
	"beegoHello1/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//①获取name、age、sex。
	userName :=c.Ctx.Input.Query("user")
	password :=c.Ctx.Input.Query("psd")
	//②做数据对比,使用固定数据进行校验
	if userName != "qiangzi" || password !="123456" {
		//代表错误
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("成功"))

	//③

	c.Data["Website"] = "www.github.com"
	c.Data["Email"] = "hhsp.app"
	c.TplName = "index.tpl"
}

//定义一个post请求 postman浏览器插件
/*func (c *MainController) Post()  {
	name :=c.Ctx.Request.FormValue("name")
	age :=c.Ctx.Request.FormValue("age")
	sex :=c.Ctx.Request.FormValue("sex")
	fmt.Println(sex)
	if name !="qiangzi" && age !="11"{
		c.Ctx.WriteString("数据错误")
		return
	}
	c.Ctx.WriteString("成功")
}

 */
/*
该方法用于
 */
func (c *MainController) post() {
	var person models.Perspn
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil {
		c.Ctx.WriteString("错误")
		return
	}
	err =json.Unmarshal(dataBytes,&person)
	if err !=nil {
		c.Ctx.WriteString("解析失败")
		return
	}
	fmt.Print("姓名",person.Name)
	fmt.Print("年龄",person.Age)
	fmt.Print("性别",person.Sex)
	c.Ctx.WriteString("成功")
}