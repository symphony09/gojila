package model

import "github.com/jinzhu/gorm"

type Panel struct {
	gorm.Model
	Name      string
	Tasks     []Task
	ProjectID uint
}
