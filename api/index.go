package api

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/model"
	"github.com/symphony09/gojila/service"
	"net/http"
)

func CreateProject(c *gin.Context) {
	var p model.Project
	_ = c.ShouldBind(&p)
	service.NewProject(&p)
}

func Projects(c *gin.Context) {
	c.JSON(http.StatusOK,
		service.GetProjects(),
	)
}
