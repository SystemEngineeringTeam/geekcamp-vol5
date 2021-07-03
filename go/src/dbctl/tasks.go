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
	/*

		for rows.Next() {
			// SQLの結果一行分を受け取る変数を宣言
			var id int
			var Detail string
			var IsAvailable int
			// 受け取る
			rows.Scan(&id, &Detail, &IsAvailable)
			// 受け取ったものをTasksに追加する．
			Tasks = append(Tasks, Task{ID: id, Name: name, Description: description, SubmitTime: submitTime, Label: label})
		}
	*/

	return Tasks, nil
}
