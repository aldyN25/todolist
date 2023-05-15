package configs

import (
	"os"
	"sync"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
	Host string
}

type DbConfig struct {
	Host        string
	Port        string
	Dbname      string
	Username    string
	Password    string
	DbIsMigrate bool
}

type Configs struct {
	Appconfig AppConfig
	Dbconfig  DbConfig
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	configs = &Configs{
		Appconfig: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Env:  os.Getenv("APP_ENV"),
			Port: os.Getenv("APP_PORT"),
			Host: os.Getenv("APP_HOST"),
		},
		Dbconfig: DbConfig{
			Host:        os.Getenv("MYSQL_HOST"),
			Port:        os.Getenv("MYSQL_PORT"),
			Dbname:      os.Getenv("MYSQL_DBNAME"),
			Username:    os.Getenv("MYSQL_USER"),
			Password:    os.Getenv("MYSQL_PASSWORD"),
			DbIsMigrate: os.Getenv("DB_IS_MIGRATE") == "true",
		},
	}

	return configs
}
