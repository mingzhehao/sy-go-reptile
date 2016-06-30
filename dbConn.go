package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/config"
)

var (
	DB_NAME,
	DB_USER,
	DB_PASS,
	DB_HOST,
	DB_PORT string
)

func init() {
	c, _ := config.ReadDefault("config/config.ini")
	DB_USER, _ = c.String("DB", "DB_USER")
	DB_PASS, _ = c.String("DB", "DB_PASS")
	DB_HOST, _ = c.String("DB", "DB_HOST")
	DB_NAME, _ = c.String("DB", "DB_NAME")
	DB_PORT, _ = c.String("DB", "DB_PORT")
	fmt.Println("\n当前服务器IP及端口为：", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
}

func GetDbConnection() (db *sql.DB, errstr error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		fmt.Println("\nconnection mysql error")
		return db, err
	}
	//defer db.Close()
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("\nopen mysql error ", err)
		return db, err
	}
	return db, nil
}
