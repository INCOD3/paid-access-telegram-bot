package storage

import "github.com/w1png/paid-access-telegram-bot/models"

type Storage interface {
  GetUser(telegramID int64) (*models.User, error)
  CreateUser(telegramID int64) (*models.User, error)

}

var CurrentStorage Storage

func InitStorage(storageType string) error {
  switch storageType {
  }

  return nil
}

