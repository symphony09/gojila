package service

import (
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/repository"
)

func CreateOneTask(t *model.Task) {
	repository.CreateOneTask(t)
}

func GetPanelTasks(pid int) {
	repository.FindTaskByPanelID(pid)
}
