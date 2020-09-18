package service

import (
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/repository"
)

func RecordEvent(e *model.Event) {
	repository.CreateEvent(e)
}

func GetTaskEvents(tid int) []*model.Event {
	return repository.GetEventsByTaskID(tid)
}
