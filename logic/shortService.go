package logic

import (
	"crypto/sha1"
	"fmt"
	"shorturl/models"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/astaxie/beego"
	"github.com/sqids/sqids-go"
)

type ShortService struct{}
type Data struct {
	Url      string
	Sha1     string
	Count    int
	CreateAt int32
	Creator  string
}

func (s *ShortService) GetShortUrl(code string) (string, error) {
	sq, error := sqids.New(sqids.Options{
		MinLength: 6,
	})
	if error != nil {
		return "加密组件启动异常", error
	}
	id := sq.Decode(code)
	if len(id) == 0 {
		return "当前编码不合法", nil
	}
	query, u := gplus.NewQuery[models.Shorturl]()
	query.Eq(&u.Id, id)
	one, _ := gplus.SelectOne(query)
	if len(one.Url) > 0 {
		//更新操作
		number := one.Count + 1
		q, link := gplus.NewQuery[models.Shorturl]()
		q.Eq(&link.Id, id).Set(&link.Count, number)

		gplus.Update(q)

		return one.Url, nil
	}
	return "链接识别异常", nil
}
func (s *ShortService) CreateShortUrl(longUrl string) (string, error) {

	h := sha1.New()
	h.Write([]byte(longUrl))
	hash := h.Sum(nil)

	query, u := gplus.NewQuery[models.Shorturl]()
	query.Eq(&u.Sha1, fmt.Sprintf("%x", hash))
	one, _ := gplus.SelectOne(query)

	sq, error := sqids.New(sqids.Options{
		MinLength: 6,
	})
	if error != nil {
		return "加密组件启动异常", error
	}
	if one.Id > 0 {
		id, _ := sq.Encode([]uint64{uint64(one.Id)}) // "86Rf07"
		return beego.AppConfig.String("domain") + id, nil
	} else {

		// ip, error := (*util.Helper).GetIpToLong
		// if error != nil {
		// 	return "获取ip异常", error
		// }
		shorturl := models.Shorturl{Url: longUrl, Sha1: fmt.Sprintf("%x", hash), CreateAt: int32(time.Now().Unix()), Creator: "20250720"}
		result := gplus.Insert(&shorturl)
		if result.Error != nil {
			return "数据库插入异常", result.Error
		}
		id, _ := sq.Encode([]uint64{uint64(shorturl.Id)}) // "86Rf07"
		return beego.AppConfig.String("domain") + id, nil
	}
}
