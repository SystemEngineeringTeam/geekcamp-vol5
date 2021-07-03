package apifuncs

import "github.com/gin-gonic/gin"

type JsonTaskGet struct {
	TaskID      int    `json:"task_id,omitempty"`
	Detail      string `json:"detail,omitempty"`
	IsAvailable bool   `json:"is_available,omitempty"`
}

func TaskGetHandler(c *gin.Context) {
	var json JsonTaskGet
	if err:=c. {

	}
}
