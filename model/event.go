package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Event struct {
	gorm.Model
	From   string
	To     string
	Note   string
	Time   time.Time
	TaskID uint
}
