package main

import (
	"github.com/SystemEngineeringTeam/geekcamp-vol5/apifuncs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/task", apifuncs.TaskGetHandler)
	r.POST("/task", apifuncs.PostTaskHandler)
	r.Run()
}
