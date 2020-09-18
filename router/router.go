package router

import (
	"github.com/gin-gonic/gin"
	"github.com/symphony09/gojila/api"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) {
	// MiddleWares.
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect route.")
	})

	kanbanAPI := g.Group("/kanban")
	{
		task := kanbanAPI.Group("/task")
		{
			task.GET("/list", api.Tasks)
			task.POST("/create", api.CreateTask)
			task.GET("/event", api.TaskEvents)
			task.POST("/event", api.Action)
		}
	}
}
