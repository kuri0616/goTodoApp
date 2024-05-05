package config

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Name string
	User string
	Pass string
	Host string
}

var MySqlConf DBConfig

func LoadDBConfig() {
	dbConfig := DBConfig{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
	}
	MySqlConf = dbConfig
}
