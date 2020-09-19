package api

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/service"
	"net/http"
	"time"
)

func CreateTask(c *gin.Context) {
	var payload model.Task
	_ = c.ShouldBind(&payload)
	service.CreateOneTask(&payload)
}

func Tasks(c *gin.Context) {
	c.JSON(http.StatusOK,
		service.GetPanelTasks,
	)
}

func Action(c *gin.Context) {
	var payload model.Event
	_ = c.ShouldBind(&payload)
	payload.Time = time.Now()
	service.RecordEvent(&payload)
	if payload.From != payload.To {
		service.ChangeTaskPanel(payload.TaskID, payload.To)
	}
}
