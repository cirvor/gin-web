package database

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	log.Println("++++++++++++++++++ mysql connecting +++++++++++++++++++")
	var err error
	Db, err = gorm.Open("mysql", "root:121837@tcp(127.0.0.1:3306)/learn?charset=utf8mb4")
	if err != nil {
		log.Fatalln(err)
	}

	// 设置最大连接数
	Db.DB().SetMaxOpenConns(30)
	// 设置最大空闲连接数
	Db.DB().SetMaxIdleConns(10)
	// 设置每个链接的过期时间
	Db.DB().SetConnMaxLifetime(time.Second * 5)

	if err := Db.DB().Ping(); err != nil {
		log.Fatalln(err)
	}
}
