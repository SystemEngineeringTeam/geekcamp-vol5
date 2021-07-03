package main

import (
	"github.com/SystemEngineeringTeam/geekcamp-vol5/apifuncs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.GET("/task", apifuncs.TaskGetHandler)
	r.POST("/task", apifuncs.PostTaskHandler)

	r.POST("/admin/task", apifuncs.RegisterNewTask)
	r.PUT("/admin/task", apifuncs.ToggleAvailable)

	r.Run(":80")
}
