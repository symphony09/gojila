package service

import (
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/repository"
)

func NewTask(t *model.Task) {
	repository.CreateOneTask(t)
}

func GetTask(tid int) *model.Task {
	return repository.FindTaskByID(tid)
}

func GetPanelTasks(pid int) []*model.Task {
	return repository.FindTaskByPanelID(pid)
}

func ChangeTaskPanel(tid uint, pn string) {
	repository.UpdateTaskPanel(tid, pn)
}

func RecordEvent(e *model.Event) {
	repository.CreateEvent(e)
}

func GetTaskEvents(tid int) []*model.Event {
	return repository.FindEventsByTaskID(tid)
}
