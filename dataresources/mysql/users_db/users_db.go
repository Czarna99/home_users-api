package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_ROOT_PASSWORD")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("MYSQL_DATABASE")
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}
