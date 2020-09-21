package repository

import (
	"github.com/symphony09/gojila/global"
	"github.com/symphony09/gojila/model"
)

func FindEventsByProjectID(pid int) []*model.Event {
	var events []*model.Event
	subQ1 := global.DB().Model(&model.Panel{}).Select("id").Where("project_id = ?", pid).SubQuery()
	subQ2 := global.DB().Model(&model.Task{}).Select("id").Where("panel_id in (?)", subQ1).SubQuery()
	global.DB().Where("task_id in (?)", subQ2).Find(&events)
	return events
}

func FindEventsByTaskID(id int) []*model.Event {
	var events []*model.Event
	global.DB().First(&model.Task{}, id).Association("Events").Find(&events)
	return events
}

func CreateEvent(e *model.Event) {
	global.DB().Create(e)
}
