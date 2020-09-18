package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Name        string
	Alias       string
	Tag         string
	Sort        int
	Description string
	Events      []Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PanelID     uint
}
