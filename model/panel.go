package model

import "github.com/jinzhu/gorm"

type Panel struct {
	gorm.Model
	Name      string
	Tasks     []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProjectID uint
}
