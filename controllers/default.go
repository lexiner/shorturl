package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"shorturl/models"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/astaxie/beego"
	"github.com/sqids/sqids-go"
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
	s, error := sqids.New(sqids.Options{
		MinLength: 6,
	})
	if error != nil {
		c.Ctx.WriteString(error.Error())
		return
	}
	id := s.Decode(code)
	if id == nil {
		c.Ctx.WriteString("Error id")
		return
	}
	query, u := gplus.NewQuery[models.Shorturl]()
	query.Eq(&u.Id, id)
	one, _ := gplus.SelectOne(query)

	number := one.Count + 1
	q, link := gplus.NewQuery[models.Shorturl]()
	q.Eq(&link.Id, id).Set(&link.Count, number)
	gplus.Update(q)

	c.Data["url"] = one.Url
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

	h := sha1.New()
	h.Write([]byte(url))
	hash := h.Sum(nil)

	query, u := gplus.NewQuery[models.Shorturl]()
	query.Eq(&u.Sha1, fmt.Sprintf("%x", hash))
	one, resultDb := gplus.SelectOne(query)
	if resultDb.Error != nil {
		fmt.Println(resultDb.Error)
		// return
	}
	if one.Id > 0 {
		s, error := sqids.New(sqids.Options{
			MinLength: 6,
		})
		if error != nil {
			fmt.Println(error.Error())
			return
		}
		id, _ := s.Encode([]uint64{uint64(one.Id)}) // "86Rf07"
		c.Data["json"] = map[string]interface{}{"data": beego.AppConfig.String("domain") + id, "msg": "success", "code": "200"}
		c.ServeJSON()
		return
	} else {
		shorturl := models.Shorturl{Url: url, Sha1: fmt.Sprintf("%x", hash), CreateAt: int32(time.Now().Unix()), Creator: "2015092122"}
		result := gplus.Insert(&shorturl)
		if result.Error != nil {
			fmt.Println(result.Error)
			c.Ctx.WriteString("Error inserting data")
			c.ServeJSON()
			return
		}
		s, _ := sqids.New(sqids.Options{
			MinLength: 6,
		})
		id, _ := s.Encode([]uint64{uint64(shorturl.Id)}) // "86Rf07"
		fmt.Println(id)
		c.Data["json"] = map[string]interface{}{"data": beego.AppConfig.String("domain") + id, "msg": "success", "code": "200"}
		c.ServeJSON()
	}

}
