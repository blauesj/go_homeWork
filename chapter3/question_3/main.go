package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB
var schema = `
CREATE TABLE accounts  (
    id int,
    balance  decimal(5,2),
);

CREATE TABLE transactions  (
    id int,
    from_account_id  int not null ,
    to_account_id  int,
    amount decimal(5,2)
)`

type Account struct {
	Id      int     `db:"id"`
	Balance float32 `db:"balance"`
}

type Transaction struct {
	Id            int     `db:"id"`
	FromAccountId int     `db:"from_account_id "`
	ToAccountId   int     `db:"to_account_id "`
	Amount        float32 `db:"amount"`
}

func initDB() (err error) {
	db, err = sqlx.Connect("mysql", "root:mysql@tcp(127.0.0.1:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local")
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
}
