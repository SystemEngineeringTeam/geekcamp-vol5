package dbctl

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

const errFormat = "%v\nfunction:%v file:%v line:%v\n"

var db *sql.DB

const host = "geekcamp_mysql:3306"

// const host = "localhost:33060"

func init() {
	var err error

	db, err = sql.Open("mysql", "user:password@tcp("+host+")/backend_db")

	if err != nil {
		pc, file, line, _ := runtime.Caller(0)
		f := runtime.FuncForPC(pc)
		log.Printf(errFormat, err, f.Name(), file, line)

		panic("Can't Open database.")
	} else {
		fmt.Println("SQL is working.")
	}

}
