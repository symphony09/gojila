package service

import (
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/repository"
)

func NewProject(p *model.Project) {
	repository.CreateOneProject(p)
}

func GetProjects() []*model.Project {
	return repository.FindProjects()
}

func GetProjectPanels(projectID int) []*model.Panel {
	return repository.FindProjectPanelsByID(projectID)
}

func NewPanel(p *model.Panel) {
	repository.CreateOnePanel(p)
}

func GetProjectTasks(pid int) []*model.Task {
	return repository.FindTasksByProjectID(pid)
}

func GetProjectEvents(pid int) []*model.Event {
	return repository.FindEventsByProjectID(pid)
}
