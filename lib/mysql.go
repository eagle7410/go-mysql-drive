package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type (
	EnvInterface interface {
		GetMysqlUser() *string
		GetMysqlPassword() *string
		GetMysqlHost() *string
		GetMysqlPort() *string
		GetMysq1lDatabase() *string
	}
)

func Init (env EnvInterface) (db *sql.DB, version string, err error) {
	host := *env.GetMysqlHost()

	if host == "" {
		host = "127.0.0.1"
	}

	port := *env.GetMysqlPort()

	if port == "" {
		port = "3306"
	}

	dbName := *env.GetMysq1lDatabase()

	if dbName == "" {
		dbName = "sys"
	}

	linkConnect := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		*env.GetMysqlUser(), *env.GetMysqlPassword(),
		host, port, dbName)


	db, err = sql.Open("mysql", linkConnect)

	if err != nil {
		return db, version, err
	}

	err = db.QueryRow("SELECT VERSION()", ).Scan(&version)

	return db, version, err
}
