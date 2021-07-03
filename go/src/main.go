package main

import (
	"fmt"

	"github.com/SystemEngineeringTeam/geekcamp-vol5/dbctl"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbctl.Test()
	fmt.Println("he..hell..helloworld!")
}
