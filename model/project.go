package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name   string
	Panels []Panel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Hooks  []Hook `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
