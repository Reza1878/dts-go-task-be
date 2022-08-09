package app

import (
	"database/sql"
	"dts-task/helper"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	DB_NAME := helper.GetDotEnvVariable("DB_NAME")
	DB_PASSWORD := helper.GetDotEnvVariable("DB_PASSWORD")
	DB_USERNAME := helper.GetDotEnvVariable("DB_USERNAME")
	DB_HOST := helper.GetDotEnvVariable("DB_HOST")
	DB_PORT := helper.GetDotEnvVariable("DB_PORT")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME))

	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
