package repository

import (
	"github.com/symphony09/gojila/global"
	"github.com/symphony09/gojila/model"
)

func CreateOneTask(t *model.Task) {
	global.DB().Create(t)
}

func FindTaskByID(id int) *model.Task {
	var t model.Task
	global.DB().First(&t, id)
	return &t
}

func FindTaskByPanelID(id int) []*model.Task {
	var tasks []*model.Task
	global.DB().First(&model.Panel{}, id).Association("Tasks").Find(&tasks)
	return tasks
}

func FindTasksByProjectID(pid int) []*model.Task {
	var tasks []*model.Task
	global.DB().Where(
		"panel_id in (?)",
		global.DB().Model(&model.Panel{}).Select("id").Where("project_id = ?", pid).SubQuery()).Find(&tasks)
	return tasks
}

func UpdateTaskPanel(tid uint, pn string) {
	var task model.Task
	var oldPanel, newPanel model.Panel
	global.DB().First(&task, tid)
	global.DB().First(&oldPanel, task.PanelID)
	global.DB().Where(&model.Panel{
		Name:      pn,
		ProjectID: oldPanel.ProjectID,
	}).First(&newPanel)
	task.PanelID = newPanel.ID
	global.DB().Save(&task)
}
