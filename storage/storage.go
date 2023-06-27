package storage

import (
	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/models"
)

type Storage interface {
  SaveChannel(channel *models.Channel) error
  GetChannelById(id string) (*models.Channel, error)

  SaveSubscription(subscription *models.Subscription) error
  GetSubscriptionById(id string) (*models.Subscription, error)
  GetSubscriptionsByChannelId(channelId string) ([]*models.Subscription, error)
  GetSubscriptionsByUserId(userId string) ([]*models.Subscription, error)

  SaveUser(user *models.User) error
  GetUserById(id int64) (*models.User, error)
}

var CurrentStorage Storage

func InitStorage(storageType string, is_test bool) error {
  switch storageType {
    case "postgres": {
      postgresStorage, err := NewPostgresStorage(is_test)
      if err != nil {
        return err
      }
      CurrentStorage = postgresStorage
    }
    default: return errors.NewStorageNotFoundError(storageType)
  }

  return nil
}

