package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	log.Println("++++++++++++++++++ mysql connecting +++++++++++++++++++")
	var err error
	Db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/learn?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, _ := Db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(30)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Second * 5)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}
}
