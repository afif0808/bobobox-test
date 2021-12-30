package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	e := echo.New()

	e.Start(":8080")
}

func initMySQLDB() (readDB, writeDB *sqlx.DB) {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, db)
	log.Println(dsn)
	readDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	err = readDB.Ping()
	if err != nil {
		log.Panic(err)
	}

	writeDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	err = readDB.Ping()
	if err != nil {
		log.Panic(err)
	}

	return
}
