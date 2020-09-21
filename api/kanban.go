package api

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/service"
	"net/http"
	"strconv"
	"time"
)

func CreatePanel(c *gin.Context) {
	var p model.Panel
	_ = c.ShouldBind(&p)
	service.NewPanel(&p)
}

func Panels(c *gin.Context) {
	pid := c.Query("project_id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK,
			service.GetProjectPanels(id),
		)
	}
}

func CreateTask(c *gin.Context) {
	var t model.Task
	_ = c.ShouldBind(&t)
	service.NewTask(&t)
}

func Task(c *gin.Context) {
	tid := c.Query("task_id")
	id, err := strconv.Atoi(tid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK,
			service.GetTask(id),
		)
	}
}

func Tasks(c *gin.Context) {
	pid := c.Query("panel_id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK,
			service.GetPanelTasks(id),
		)
	}
}

func Action(c *gin.Context) {
	var e model.Event
	_ = c.ShouldBind(&e)
	e.Time = time.Now()
	service.RecordEvent(&e)
	if e.From != e.To {
		service.ChangeTaskPanel(e.TaskID, e.To)
	}
}

func TaskEvents(c *gin.Context) {
	taskId := c.Query("task_id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK, service.GetTaskEvents(id))
	}
}
