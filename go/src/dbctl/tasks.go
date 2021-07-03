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

	rows, err := db.Query("SELECT id, detail, isAvailable from tasks;")
	if err != nil {
		fmt.Println("SQL QUERY ERR")
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var Tasks []Task

	for rows.Next() {

		var id int
		var Detail string
		var isAvailable int

		tmp := false

		rows.Scan(&id, &Detail, &isAvailable)
		if isAvailable == 0 {
			tmp = false

		} else if isAvailable == 1 {
			tmp = true
		}
		Tasks = append(Tasks, Task{TaskID: id, Detail: Detail, IsAvailable: tmp})
	}
	return Tasks, nil
}
