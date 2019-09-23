package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"translate-svc/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	//host := env.LoadOrDefault("DB_HOST", "127.0.0.1")
	//port := env.LoadOrDefault("DB_PORT", "3306")
	//user := env.LoadOrDefault("DB_USER", "root")
	//password := env.LoadOrDefault("DB_PASSWORD", "root")
	//dbname := env.LoadOrDefault("DB_NAME", "iam")
	host := config.DbHost
	port := config.DbPort
	user := config.DbUser
	password := config.DbPassword
	dbname := config.DbName

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=True&charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	DB, err = gorm.Open("mysql", connection)

	if err != nil {
		log.Panic(err)
	}

	// Important.
	DB.DB().SetMaxOpenConns(config.DbMaxOpenConnections)
	DB.DB().SetMaxIdleConns(config.DbMaxIdleConnections)

	debug, _ := strconv.ParseBool(config.DbDebug)
	DB.LogMode(debug)
}
