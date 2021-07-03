package apifuncs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonTaskGet struct {
	TaskID      int    `json:"task_id"`
	Detail      string `json:"detail"`
	IsAvailable bool   `json:"is_available"`
}

type JsonTaskPost struct {
	TaskID int `json:"task_id"`
}

func TaskGetHandler(c *gin.Context) {
	var res []JsonTaskGet

	// TODO: データベースの関数呼び出し

	res = append(res, JsonTaskGet{9999, "hoge", true})

	c.JSON(http.StatusOK, res)
}

func PostTaskHandler(c *gin.Context) {
	var req JsonTaskPost
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse JSON"})
		return
	}
	// todo: データベースの関数呼び出し

}
