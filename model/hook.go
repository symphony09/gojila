package model

import "github.com/jinzhu/gorm"

type Hook struct {
	gorm.Model
	Api       string
	Payloads  []Payload
	ProjectID uint
}

type Payload struct {
	gorm.Model
	Key    string
	Value  string
	HookID uint
}
