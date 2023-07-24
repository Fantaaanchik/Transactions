package service

import (
	"transactions/internal/db"
	"transactions/models"
)

type Services struct {
	Repository RepositoryInterface
}

func NewService(repo RepositoryInterface) *Services {
	return &Services{Repository: repo}
}

type RepositoryInterface interface {
	UpdateUserBalance(users models.User) error
	GetUserByPhone(phone string) (models.User, error)
	CreateTransaction(transactions models.Transaction) error
}

func (s *Services) MakeTransaction(senderPhone string, receiverPhone string, amount int) error {

	tx := db.DB.Begin()
	var phoneForTransaction models.Transaction
	sender, err := s.Repository.GetUserByPhone(senderPhone)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	receiver, err := s.Repository.GetUserByPhone(receiverPhone)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if sender.Balance < amount {
		_ = tx.Rollback()
		return err
	}

	phoneForTransaction.SenderPhone = sender.Number

	phoneForTransaction.ReceiverPhone = receiver.Number

	transaction := models.Transaction{
		SenderPhone:   phoneForTransaction.SenderPhone,
		ReceiverPhone: phoneForTransaction.ReceiverPhone,
		Amount:        amount,
	}

	err = s.Repository.CreateTransaction(transaction)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()

	return nil
}
