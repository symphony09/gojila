package api

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/service"
	"net/http"
	"strconv"
)

func TaskEvents(c *gin.Context) {
	taskId := c.Query("task_id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK, service.GetTaskEvents(id))
	}
}