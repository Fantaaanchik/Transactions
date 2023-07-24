package repository

import (
	"gorm.io/gorm"
	"log"
	"transactions/internal/db"
	"transactions/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type UserRepository struct {
	UserRep models.User
}

func (r *Repository) UpdateUserBalance(users models.User) error {

	err := db.DB.Create(&users).Error
	if err != nil {
		log.Println("Cannot get user from DB, err: ", err.Error())
	}
	return nil
}

func (r *Repository) GetUserByPhone(phone string) (models.User, error) {
	var users models.User
	err := db.DB.Where("number = ?", phone).Find(&users).Error
	if err != nil {
		log.Println("Cannot get user from DB, err: ", err.Error())
	}

	return users, nil
}

func (r *Repository) CreateTransaction(transactions models.Transaction) error {

	err := db.DB.Create(&transactions).Error
	if err != nil {
		log.Println("cannot create user transaction (CreateTransaction), err: ", err.Error())
	}

	amount := transactions.Amount
	receiverUserNumber := transactions.ReceiverPhone
	SenderUserNumber := transactions.SenderPhone

	err = db.DB.Model(&models.User{}).Where("number = ?", SenderUserNumber).Update("user_balance", gorm.Expr("user_balance - ?", amount)).Error
	if err != nil {
		log.Println("cannot subtract amount from user1, err:", err.Error())
	}

	err = db.DB.Model(&models.User{}).Where("number = ?", receiverUserNumber).Update("user_balance", gorm.Expr("user_balance + ?", amount)).Error
	if err != nil {
		log.Println("cannot add amount to user2, err:", err.Error())
	}

	return nil
}
