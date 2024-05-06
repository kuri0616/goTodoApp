package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Name string
	User string
	Pass string
	Host string
}

var MySqlConf DBConfig

func loadDBConfig() {
	dbConfig := DBConfig{
		Name: os.Getenv("DBNAME"),
		User: os.Getenv("DBUSER"),
		Pass: os.Getenv("DBPASS"),
		Host: os.Getenv("DBHOST"),
	}
	MySqlConf = dbConfig
}
func InitDB() *sqlx.DB {
	loadDBConfig()
	dbConf := fmt.Sprintf("%s:%s@tcp(%s)/%s", MySqlConf.User, MySqlConf.Pass, MySqlConf.Host, MySqlConf.Name)
	log.Println(dbConf)
	db, err := sqlx.Connect("mysql", dbConf)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
