package dbctl

import (
	"fmt"
	"log"
)

type Task struct {
	TaskID      int    `json:"task_id"`
	Detail      string `json:"detail"`
	IsAvailable bool   `json:"is_available"`
}

//Test はテストする関数
func Test() {
	fmt.Println("Test関数")
}

func Gettasks() ([]Task, error) {

	rows, err := db.Query("select id, detail, isAvaliable, from tasks")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var Tasks []Task
	return Tasks, nil
}
