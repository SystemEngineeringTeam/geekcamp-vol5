package apifuncs

import (
	"math/rand"
	"net/http"

	"github.com/SystemEngineeringTeam/geekcamp-vol5/dbctl"
	"github.com/gin-gonic/gin"
)

type JsonTaskPost struct {
	TaskID int `json:"task_id"`
}

func TaskGetHandler(c *gin.Context) {

	// TODO: データベースの関数呼び出し
	json, err := dbctl.Gettasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	res := getThreeTasksRandomly(json)

	c.JSON(http.StatusOK, res)
}

func getThreeTasksRandomly(tasks []dbctl.Task) []dbctl.Task {
	three := make([]dbctl.Task, 0)
	numbers := make([]int, 0)
	isUnique := true

	for i := 0; len(three) != 3; i++ {
		isUnique = true
		n := rand.Intn(len(tasks))
		for _, v := range numbers {
			if n == v {
				isUnique = false
			}
		}
		if isUnique {
			numbers = append(numbers, n)
			three = append(three, tasks[n])
		}
	}
	return three
}

func PostTaskHandler(c *gin.Context) {
	var req JsonTaskPost
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse JSON"})
		return
	}
	// todo: データベースの関数呼び出し
	if err := dbctl.InsertCount(req.TaskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error"})
		return
	}
	c.Status(http.StatusOK)
}
