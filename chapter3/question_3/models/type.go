package models

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

type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}

type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

/**
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
*/
