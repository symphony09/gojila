package api

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/service"
	"net/http"
	"strconv"
)

func ProjectTasks(c *gin.Context) {
	pid := c.Query("project_id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK, service.GetProjectTasks(id))
	}
}

func ProjectEvents(c *gin.Context) {
	pid := c.Query("project_id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK, service.GetProjectEvents(id))
	}
}
