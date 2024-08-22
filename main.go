package main

import (
	"log"
	_ "shorturl/routers"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDb *gorm.DB

func initDb() {
	db_user := beego.AppConfig.String("db_user")
	db_pwd := beego.AppConfig.String("db_pwd")
	db_database := beego.AppConfig.String("db_database")
	db_host := beego.AppConfig.String("db_host")
	dsn := db_user + ":" + db_pwd + "@tcp(" + db_host + ":3306)/" + db_database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
	}
	// 初始化gplus
	gplus.Init(gormDb)
}
func main() {
	initDb()
	beego.Run()
}
