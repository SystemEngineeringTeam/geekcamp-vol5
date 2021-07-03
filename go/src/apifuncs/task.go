package apifuncs

import "github.com/gin-gonic/gin"

type JsonTaskGet struct {
	TaskID      int
	Detail      string
	IsAvailable bool
}

func TaskGetHandler(c *gin.Context) {

}
