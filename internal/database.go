package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	cfg "github.com/bassamtiano/gotravel/configs/env"
	_ "github.com/go-sql-driver/mysql"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
}

type Database struct {
}

func InitDatabase() *sql.DB {
	host := config.GetString(`database.host`)
	port := config.GetString(`database.port`)
	user := config.GetString(`database.user`)
	password := config.GetString(`database.pass`)
	db_name := config.GetString(`database.name`)

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db_name)

	fmt.Println(connectionString)

	var err error
	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	return DB

}
