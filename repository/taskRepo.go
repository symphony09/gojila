package repository

import (
	"github.com/symphony09/gojila/global"
	"github.com/symphony09/gojila/model"
)

func CreateOneTask(t *model.Task) {
	global.DB().Create(t)
}

func FindTaskByPanelID(id int) []*model.Task {
	var tasks []*model.Task
	global.DB().First(&model.Panel{}, id).Association("Tasks").Find(&tasks)
	return tasks
}