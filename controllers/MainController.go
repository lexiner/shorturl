package controllers

import (
	"encoding/json"
	"shorturl/logic"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Param struct {
	LongUrl string `json:"longurl"`
}

func (c *MainController) Get() {
	c.Data["Website"] = "Shorturl By Golang Develope @tianshu"
	c.TplName = "index.tpl"
}

// 访问短链接
func (c *MainController) TargetUrl() {
	code := c.Ctx.Input.Param(":code")
	if code == "" {
		c.Ctx.WriteString("Error decoding JSON")
		return
	}
	url, error := (&logic.ShortService{}).GetShortUrl(code)
	if url == "当前编码不合法" {
		c.Ctx.WriteString(url)
		c.ServeJSON()
		return
	}
	if error != nil {
		c.Ctx.WriteString(url)
		c.ServeJSON()
		return
	}
	c.Data["url"] = url
	c.TplName = "target.tpl"
}

// 创建短链接
func (c *MainController) CreateUrl() {
	var param Param
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		c.Ctx.WriteString("Error decoding JSON")
		return
	}
	url := param.LongUrl
	url, error := (&logic.ShortService{}).CreateShortUrl(url)
	if error != nil {
		c.Ctx.WriteString(url)
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"data": url, "msg": "success", "code": "200"}
	c.ServeJSON()
}
