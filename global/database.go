package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/symphony09/gojila/model"
	"sync"
)

var db *gorm.DB

func DB() *gorm.DB {
	var once sync.Once
	once.Do(initDB)
	return db
}

func initDB() {
	c := Config.Sub("db")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.GetString("user"),
		c.GetString("password"),
		c.GetString("host"),
		c.GetString("port"),
		c.GetString("name"),
	)
	db, _ = gorm.Open("mysql", args)
	db.LogMode(true)
	db.AutoMigrate(model.Project{},model.Panel{},model.Task{},model.Event{},model.Hook{},model.Payload{})
}

func CloseDB() {
	_ = db.Close()
}
