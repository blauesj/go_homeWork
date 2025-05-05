package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go_homeWork/chapter3/question_4/models"
	"log"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(10)
	return err
}
func main() {
	err := initDB()
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	sql := `select * from books where price > ?`
	var books []models.Book
	db.Select(&books, sql, 50)
	fmt.Printf("hightest salary employee:%#v\n", books)
}
