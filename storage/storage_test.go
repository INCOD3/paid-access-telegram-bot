package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func TestNewStorage_Postgres(t *testing.T) {
  err := utils.InitConfig()
  assert.Nil(t, err)
  
  err = InitStorage("postgres", false)
  assert.Nil(t, err)
  assert.Equal(t, reflect.TypeOf(&PostgresStorage{}), reflect.TypeOf(CurrentStorage), "Type of CurrentStorage is incorrect")
}

func TestNewStorage_StorageNotFoundError(t *testing.T) {
  err := InitStorage("invalid_storage", false)
  assert.Equal(t, reflect.TypeOf(&errors.StorageNotFoundError{}), reflect.TypeOf(err), "Type of err is incorrect")
  assert.Equal(t, "Storage not found: invalid_storage", err.Error(), "Error message is incorrect")
}
