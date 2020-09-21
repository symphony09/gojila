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

	indexAPI := g.Group("/index")
	{
		project := indexAPI.Group("/project")
		{
			project.POST("/create", api.CreateProject)
			project.GET("/list", api.Projects)
		}
	}

	kanbanAPI := g.Group("/kanban")
	{
		panel := kanbanAPI.Group("/panel")
		{
			panel.POST("/create", api.CreatePanel)
			panel.GET("/list", api.Panels)
		}

		task := kanbanAPI.Group("/task")
		{
			task.POST("/create", api.CreateTask)
			task.GET("/info", api.Task)
			task.GET("/list", api.Tasks)
			task.POST("/event", api.Action)
			task.GET("/event", api.TaskEvents)
		}
	}

	overviewAPI := g.Group("/overview")
	{
		overviewAPI.GET("/task", api.ProjectTasks)
		overviewAPI.GET("/event", api.ProjectEvents)
	}
}
