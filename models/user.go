package models

import (
	"github.com/w1png/paid-access-telegram-bot/utils"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model

  TelegramID int64 `gorm:"unique"`
  Role utils.Role
  Balance int64

  Subscriptions []Subscription `gorm:"many2many:user_subscriptions;"`
}

func NewUser(telegramID int64) *User {
  return &User{
    TelegramID: telegramID,
    Role: utils.User,
    Balance: 0,
  }
}

func (u *User) IsAdmin() bool {
  return u.Role == utils.Admin
}

func (u *User) AddBalance(amount int64) {
  u.Balance += amount
}

