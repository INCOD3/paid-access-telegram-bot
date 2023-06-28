package storage

import (
	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/models"
)

type Storage interface {
  SaveChannel(channel *models.Channel) error
  GetChannelByName(name string) (*models.Channel, error)
  DeleteChannelByName(name string) error
  GetChannels() ([]*models.Channel, error)

  SaveSubscription(subscription *models.Subscription) error
  GetSubscriptionById(id string) (*models.Subscription, error)
  GetSubscriptionsByChannelName(channelName string) ([]*models.Subscription, error)
  GetSubscriptionsByTelegramId(telegramId int64) ([]*models.Subscription, error)

  SaveUser(user *models.User) error
  GetUserByTelegramId(telegramId int64) (*models.User, error)
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

