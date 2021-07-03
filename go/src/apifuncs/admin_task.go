package apifuncs

import (
	"net/http"

	"github.com/SystemEngineeringTeam/geekcamp-vol5/dbctl"
	"github.com/gin-gonic/gin"
)

func ToggleAvailable(c *gin.Context) {
	var json IdOnlyStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, JsonParsingError)
		return
	}

	if err := dbctl.IsAvailableReverse(json.TaskID); err != nil {
		c.JSON(http.StatusInternalServerError, DatabaseError)
		return
	}
	c.Status(http.StatusOK)
}

type NewTask struct {
	Detail      string `json:"detail,omitempty"`
	IsAvailable bool   `json:"is_available,omitempty"`
}

func RegisterNewTask(c *gin.Context) {
	var json NewTask
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, JsonParsingError)
		return
	}

	if err := dbctl.InsertTask(json.Detail, json.IsAvailable); err != nil {
		c.JSON(http.StatusInternalServerError, DatabaseError)
		return
	}

	c.Status(http.StatusOK)
}
