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
  db *gorm.DB
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
  err := s.db.AutoMigrate(&models.Channel{})
  if err != nil {
    return err
  }
  err = s.db.AutoMigrate(&models.Subscription{})
  if err != nil {
    return err
  }
  err = s.db.AutoMigrate(&models.User{})
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
  
  s.db = db
  return nil
}

func (s *PostgresStorage) SaveChannel(channel *models.Channel) error {
  result := s.db.Create(channel)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (s *PostgresStorage) GetChannelById(id string) (*models.Channel, error) {
  var channel models.Channel
  result := s.db.Where("id = ?", id).First(&channel)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return &channel, errors.NewObjectNotFoundError(fmt.Sprintf("channel with id %s", id))
    }
    return nil, result.Error
  }
  return &channel, nil
}

func (s *PostgresStorage) SaveSubscription(subscription *models.Subscription) error {
  result := s.db.Create(subscription)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (s *PostgresStorage) GetSubscriptionById(id string) (*models.Subscription, error) {
  var subscription models.Subscription
  result := s.db.Where("id = ?", id).First(&subscription)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return &subscription, errors.NewObjectNotFoundError(fmt.Sprintf("subscription with id %s", id))
    }
    return nil, result.Error
  }
  return &subscription, nil
}

func (s *PostgresStorage) GetSubscriptionsByChannelId(channelId string) ([]*models.Subscription, error) {
  var subscriptions []*models.Subscription
  result := s.db.Where("channel_id = ?", channelId).Find(&subscriptions)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return subscriptions, errors.NewObjectNotFoundError(fmt.Sprintf("subscriptions with channel id %s", channelId))
    }
    return nil, result.Error
  }
  return subscriptions, nil
}

func (s *PostgresStorage) SaveUser(user *models.User) error {
  result := s.db.Save(user)
  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (s *PostgresStorage) GetUserById(id int64) (*models.User, error) {
  var user models.User
  result := s.db.Where("id = ?", id).First(&user)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return &user, errors.NewObjectNotFoundError(fmt.Sprintf("user with id %s", strconv.FormatInt(id, 10)))
    }
    return nil, result.Error
  }
  return &user, nil
}

func (s *PostgresStorage) GetSubscriptionsByUserId(userId string) ([]*models.Subscription, error) {
  var subscriptions []*models.Subscription
  result := s.db.Where("user_id = ?", userId).Find(&subscriptions)
  if result.Error != nil {
    if result.Error == gorm.ErrRecordNotFound {
      return subscriptions, errors.NewObjectNotFoundError(fmt.Sprintf("subscriptions for user with id %s", userId))
    }
    return nil, result.Error
  }
  return subscriptions, nil
}

