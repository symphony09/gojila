package repository

import (
	"github.com/symphony09/gojila/global"
	"github.com/symphony09/gojila/model"
)

func GetEventsByTaskID(id int) []*model.Event {
	var events []*model.Event
	global.DB().First(&model.Task{}, id).Association("Events").Find(&events)
	return events
}

func CreateEvent(e *model.Event) {
	global.DB().Create(e)
}
