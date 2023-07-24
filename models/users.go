package models

import "time"

type User struct {
	ID        int        `json:"user_id" gorm:"user_id"`
	FullName  string     `json:"full_name" gorm:"full_name"`
	Number    string     `json:"number" gorm:"number"`
	Balance   int        `json:"user_balance" gorm:"column:user_balance"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (u User) TableName() string {
	return "users"
}
