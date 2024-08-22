package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Shorturl struct {
	Url      string `orm:"size(255)"`
	Sha1     string `orm:"size(255)"`
	Id       int    `orm:"pk;auto"`
	Count    int    `orm:"int(10)"`
	CreateAt int32  `orm:"int32(10)"`
	Creator  string `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(Shorturl))
}

func (s *Shorturl) TableName() string {
	return "qm_short_urls"
}
