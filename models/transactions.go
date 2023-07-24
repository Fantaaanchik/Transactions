package models

import "time"

type Transaction struct {
	ID            int        `json:"id" gorm:"id"`
	SenderPhone   string     `json:"sender_phone" gorm:"sender_phone"`
	ReceiverPhone string     `json:"receiver_phone" gorm:"receiver_phone"`
	Amount        int        `json:"amount" gorm:"amount"`
	CreatedAt     *time.Time `json:"created_at" gorm:"autoCreateTime"`
}
