package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name   string
	Panels []Panel
	Hooks  []Hook
}
