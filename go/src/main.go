package main

import (
	"github.com/SystemEngineeringTeam/geekcamp-vol5/apifuncs"
	"github.com/SystemEngineeringTeam/geekcamp-vol5/dbctl"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	dbctl.Test()
	dbctl.Gettasks()
	r.GET("/task", apifuncs.TaskGetHandler)
	r.POST("/task", apifuncs.PostTaskHandler)
	r.Run(":80")
}
