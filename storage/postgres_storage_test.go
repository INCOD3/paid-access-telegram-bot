package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/models"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func TestNewPostgresStorage(t *testing.T) {
  err := utils.InitConfig()
  assert.Nil(t, err)

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)
  assert.NotNil(t, s)
}

func TestPostgresStorage_SaveGetUser(t *testing.T) {
  u := models.NewUser(123)
  u.Balance = 100

  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  err = s.SaveUser(u)
  assert.Nil(t, err)

  u2, err := s.GetUserById(u.TelegramID)
  assert.Nil(t, err)

  assert.Equal(t, u.ID, u2.ID)
  assert.Equal(t, u.Balance, u2.Balance)
}

func TestPostgresStorage_UserNotFoundError(t *testing.T) {
  s, err := NewPostgresStorage(true)
  assert.Nil(t, err)

  _, err = s.GetUserById(1)
  assert.NotNil(t, err)
  assert.Equal(t, reflect.TypeOf(&errors.ObjectNotFoundError{}), reflect.TypeOf(err))
}


