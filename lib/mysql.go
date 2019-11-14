package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type (
	EnvInterface interface {
		GetUser() *string
		GetPassword() *string
		GetHost() *string
		GetPort() *string
		GetDatabase() *string
	}
)

func Init (env EnvInterface) (db *sql.DB, version string, err error) {
	host := *env.GetHost()

	if host == "" {
		host = "127.0.0.1"
	}

	port := *env.GetPort()

	if port == "" {
		port = "3306"
	}

	dbName := *env.GetDatabase()

	if dbName == "" {
		dbName = "sys"
	}

	linkConnect := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		*env.GetUser(), *env.GetPassword(),
		host, port, dbName)


	db, err = sql.Open("mysql", linkConnect)

	if err != nil {
		return db, version, err
	}

	err = db.QueryRow("SELECT VERSION()", ).Scan(&version)

	return db, version, err
}
