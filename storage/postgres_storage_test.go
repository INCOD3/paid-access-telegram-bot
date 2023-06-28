package storage

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/models"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func setup() {
  err := utils.InitConfig()
  if err != nil {
    log.Fatal(err)
  }
}

func teardown() {
  s, err := NewPostgresStorage(true)
  if err != nil {
    log.Fatal(err)
  }

  db := s.(*PostgresStorage).DB

  err = db.Migrator().DropTable(&models.Subscription{})
  if err != nil {
    log.Fatal(err)
  }

  err = db.Migrator().DropTable(&models.User{})
  if err != nil {
    log.Fatal(err)
  }

  err = db.Migrator().DropTable(&models.Channel{})
  if err != nil {
    log.Fatal(err)
  }
}

func TestNewPostgresStorage(t *testing.T) {
  setup()

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)
  assert.NotNil(t, s)

  teardown()
}

func TestPostgresStorage_SaveGetUpdateUser(t *testing.T) {
  setup()
  u := models.NewUser(111)
  u.Balance = 100

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  err = s.SaveUser(u)
  assert.Nil(t, err)

  u2, err := s.GetUserByTelegramId(u.TelegramID)
  assert.Nil(t, err)

  assert.Equal(t, u.ID, u2.ID)
  assert.Equal(t, u.Balance, u2.Balance)

  u2.Balance = 200
  err = s.SaveUser(u2)
  assert.Nil(t, err)

  u3, err := s.GetUserByTelegramId(u.TelegramID)
  assert.Nil(t, err)
  assert.Equal(t, u2.Balance, u3.Balance)

  teardown()
}

func TestPostgresStorage_UserNotFoundError(t *testing.T) {
  setup()

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  _, err = s.GetUserByTelegramId(1)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&errors.ObjectNotFoundError{}), reflect.TypeOf(err))

  teardown()
}

func TestSaveGetChannel(t *testing.T) {
  setup()

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  c := models.NewChannel("@test", "Test channel", "Description", 100, false, 0)
  err = s.SaveChannel(c)
  assert.Nil(t, err)

  c2, err := s.GetChannelByName(c.Name)
  assert.Nil(t, err)

  assert.Equal(t, c, c2)

  teardown()
}

func TestGetChannel_NotFoundError(t *testing.T) {
  setup()

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  _, err = s.GetChannelByName("test")
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&errors.ObjectNotFoundError{}), reflect.TypeOf(err))

  teardown()
}

func TestDeleteChannel(t *testing.T) {
  setup()

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  c := models.NewChannel("test", "Test channel", "Description", 100, false, 0)
  err = s.SaveChannel(c)
  assert.Nil(t, err)

  err = s.DeleteChannelByName(c.Name)
  assert.Nil(t, err)

  _, err = s.GetChannelByName(c.Name)
  assert.NotNil(t, err)

  teardown()
}

