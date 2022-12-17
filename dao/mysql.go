package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

const (
	Host     = "127.0.0.1"
	Port     = 3306
	User     = "root"
	PassWord = "123456"
	DbName   = "todos"
)

func InitMySQL() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, PassWord, Host, Port, DbName)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping() //DB.DB().Ping()测试DB的连通性 如果DB连接成功 就返回nil
}

func Close() {
	DB.Close()
}
