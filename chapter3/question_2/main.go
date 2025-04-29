package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Account struct {
	Id      int
	Balance float32
}

type Transaction struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        float32
}

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	db.Transaction(func(tx *gorm.DB) error {
		var account = Account{Id: 1}
		var err error
		if err = tx.First(&account).Error; err != nil {
			return err
		}
		if account.Balance < 100 {
			log.Println("账号金额不足")
			return errors.New("账号金额不足")
		}
		if err = tx.Model(&Account{}).Where("id = ?", 1).Update("balance", gorm.Expr("balance - ?", 100)).Error; err != nil {
			return err
		}
		if err = tx.Model(&Account{}).Where("id = ?", 2).Update("balance", gorm.Expr("balance + ?", 100)).Error; err != nil {
			return err
		}

		if err = tx.Create(&Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 100}).Error; err != nil {
			return err
		}

		return nil
	})

}
