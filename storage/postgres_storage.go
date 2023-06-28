package storage

import (
	"fmt"
	"strconv"

	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/models"
	"github.com/w1png/paid-access-telegram-bot/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresStorage struct {
  DB *gorm.DB
}

func NewPostgresStorage(is_test bool) (Storage, error) {
  s := &PostgresStorage{}
  dbname := utils.ConfigInstance.PostgresDB
  if is_test {
    dbname = utils.ConfigInstance.PostgresTestDB
  }
  dsn := s.getDSN(
    utils.ConfigInstance.PostgresHost,
    utils.ConfigInstance.PostgresUser,
    utils.ConfigInstance.PostgresPassword,
    dbname,
    utils.ConfigInstance.PostgresPort,
  )
  err := s.initDB(dsn)
  if err != nil {
    return nil, err
  }

  err = s.autoMigrate()
  if err != nil {
    return nil, err
  }

  return s, nil
}

func (s *PostgresStorage) autoMigrate() error {
  err := s.DB.AutoMigrate(&models.Channel{})
  if err != nil {
    return err
  }
  err = s.DB.AutoMigrate(&models.Subscription{})
  if err != nil {
    return err
  }
  err = s.DB.AutoMigrate(&models.User{})
  if err != nil {
    return err
  }
  return nil
}

func (s *PostgresStorage) getDSN(host, user, password, dbname, port string) string {
  return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}

func (s *PostgresStorage) initDB(dsn string) error {
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Silent),
  })
  if err != nil {
    return err
  }
  
  s.DB = db
  return nil
}

func (s *PostgresStorage) SaveChannel(channel *models.Channel) error {
  result := s.DB.Save(channel)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (s *PostgresStorage) GetChannelByName(name string) (*models.Channel, error) {
  var channel models.Channel
  result := s.DB.Where("name = ?", name).First(&channel)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return &channel, errors.NewObjectNotFoundError(fmt.Sprintf("channel with name %s", name))
    }
    return nil, result.Error
  }
  return &channel, nil
}

func (s *PostgresStorage) DeleteChannelByName(name string) (error) {
  result := s.DB.Where("name = ?", name).Delete(&models.Channel{})
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (s *PostgresStorage) GetChannels() ([]*models.Channel, error) {
  var channels []*models.Channel
  result := s.DB.Find(&channels)
  if result.Error != nil {
    return nil, result.Error
  }

  return channels, nil
}

func (s *PostgresStorage) SaveSubscription(subscription *models.Subscription) error {
  err := s.DB.Save(subscription).Error
  if err != nil {
    if err == gorm.ErrRecordNotFound {
      return errors.NewObjectNotFoundError(fmt.Sprintf("subscription with id %d", subscription.ID))
    }
    return err
  }
  return nil
}

func (s *PostgresStorage) GetSubscriptionById(id string) (*models.Subscription, error) {
  var subscription models.Subscription
  result := s.DB.Where("id = ?", id).First(&subscription)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return &subscription, errors.NewObjectNotFoundError(fmt.Sprintf("subscription with id %s", id))
    }
    return nil, result.Error
  }
  return &subscription, nil
}

func (s *PostgresStorage) GetSubscriptionsByChannelName(channelName string) ([]*models.Subscription, error) {
  var subscriptions []*models.Subscription
  result := s.DB.Where("channel_id = ?", channelName).Find(&subscriptions)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return subscriptions, errors.NewObjectNotFoundError(fmt.Sprintf("subscriptions with channel id %s", channelName))
    }
    return nil, result.Error
  }
  return subscriptions, nil
}

func (s *PostgresStorage) SaveUser(user *models.User) error {
  result := s.DB.Save(user)
  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (s *PostgresStorage) GetUserByTelegramId(telegramId int64) (*models.User, error) {
  var user models.User
  result := s.DB.Where("telegram_id = ?", telegramId).First(&user)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return &user, errors.NewObjectNotFoundError(fmt.Sprintf("user with id %s", strconv.FormatInt(telegramId, 10)))
    }
    return nil, result.Error
  }
  return &user, nil
}

func (s *PostgresStorage) GetSubscriptionsByTelegramId(telegramId int64) ([]*models.Subscription, error) {
  var subscriptions []*models.Subscription
  result := s.DB.Where("user_id = ?", telegramId).Find(&subscriptions)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return subscriptions, errors.NewObjectNotFoundError(fmt.Sprintf("subscriptions for user with id %s", strconv.FormatInt(telegramId, 10)))
    }
    return nil, result.Error
  }
  return subscriptions, nil
}

