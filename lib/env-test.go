package lib

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"path"
	"reflect"
)

type Env struct {
	WorkDir string
	DbPort,
	DbHost,
	DbPass,
	DbName,
	DbUser string
}

func (i *Env) GetMysq1lDatabase() *string  {
	return &i.DbName
}

func (i *Env) GetMysqlPort() *string  {
	return &i.DbPort
}
func (i *Env) GetMysqlHost() *string  {
	return &i.DbHost
}
func (i *Env) GetMysqlPassword() *string  {
	return &i.DbPass
}
func (i *Env) GetMysqlUser() *string  {
	return &i.DbUser
}

func (i *Env) Init () (err error) {
	if i.WorkDir == "" {
		pwd, err := os.Getwd()

		if err != nil {
			return err
		}

		i.WorkDir = pwd
	}

	envPath := path.Join(i.WorkDir, ".env")

	if _, err := os.Stat(envPath); err == nil {
		if err := godotenv.Load(envPath); err != nil {
			return err
		}
	}

	props := map[string]bool{
		"DbPort": false,
		"DbHost": false,
		"DbName": false,
		"DbPass":    true,
		"DbUser":    true,
	}

	for prop, isRequired := range props {

		v := os.Getenv(prop)

		if isRequired == true && v == "" {
			return errors.New("Bad " + prop)
		}

		reflect.ValueOf(i).Elem().FieldByName(prop).SetString(v)
	}

	return nil
}
