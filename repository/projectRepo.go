package repository

import (
	"github.com/symphony09/gojila/global"
	"github.com/symphony09/gojila/model"
)

func CreateOneProject(p *model.Project) {
	global.DB().Create(p)
}

func FindProjects() []*model.Project {
	var projects []*model.Project
	global.DB().Find(&projects)
	return projects
}

func FindProjectPanelsByID(id int) []*model.Panel {
	var panels []*model.Panel
	global.DB().First(&model.Project{}, id).Association("Panels").Find(&panels)
	return panels
}

func CreateOnePanel(p *model.Panel) {
	global.DB().Create(p)
}
