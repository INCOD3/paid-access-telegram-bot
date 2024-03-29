package utils

import (
	"os"
	"strconv"

	"github.com/w1png/paid-access-telegram-bot/errors"
)

var ConfigInstance Config

type Config struct {
  TelegramToken string
  Language string
  StorageType string
  LoggerType string

  PostgresHost string
  PostgresPort string
  PostgresUser string
  PostgresPassword string
  PostgresDB string
  PostgresTestDB string

  MainAdminId int64
}

func InitConfig() error {
  err := ConfigInstance.GatherVariables()
  if err != nil {
    return err
  }

  return nil
}

func (c *Config) GatherVariables() error {
  var ok bool
  c.TelegramToken, ok = os.LookupEnv("TELEGRAM_TOKEN")
  if !ok {
    return errors.NewEnvironmentVariableError("TELEGRAM_TOKEN")
  }

  c.LoggerType, ok = os.LookupEnv("LOGGER_TYPE")
  if !ok {
    c.LoggerType = ""
  }

  c.Language, ok = os.LookupEnv("LANGUAGE")
  if !ok {
    return errors.NewEnvironmentVariableError("LANGUAGE")
  }

  c.StorageType, ok = os.LookupEnv("STORAGE_TYPE")
  if !ok {
    return errors.NewEnvironmentVariableError("STORAGE_TYPE")
  }

  c.PostgresHost, ok = os.LookupEnv("POSTGRES_HOST")
  if !ok {
    return errors.NewEnvironmentVariableError("POSTGRES_HOST")
  }

  c.PostgresPort, ok = os.LookupEnv("POSTGRES_PORT")
  if !ok {
    return errors.NewEnvironmentVariableError("POSTGRES_PORT")
  }

  c.PostgresUser, ok = os.LookupEnv("POSTGRES_USER")
  if !ok {
    return errors.NewEnvironmentVariableError("POSTGRES_USER")
  }

  c.PostgresPassword, ok = os.LookupEnv("POSTGRES_PASSWORD")
  if !ok {
    return errors.NewEnvironmentVariableError("POSTGRES_PASSWORD")
  }

  c.PostgresDB, ok = os.LookupEnv("POSTGRES_DB")
  if !ok {
    return errors.NewEnvironmentVariableError("POSTGRES_DB")
  }

  c.PostgresTestDB, ok = os.LookupEnv("POSTGRES_TEST_DB")
  if !ok {
    return errors.NewEnvironmentVariableError("POSTGRES_TEST_DB")
  }

  mainAdminIdRaw, ok := os.LookupEnv("MAIN_ADMIN")
  if !ok {
    return errors.NewEnvironmentVariableError("MAIN_ADMIN")
  }
  mainAdminId, err := strconv.ParseInt(mainAdminIdRaw, 10, 64)
  if err != nil {
    return errors.NewEnvironmentVariableError("MAIN_ADMIN")
  }
  c.MainAdminId = mainAdminId

  return nil
}
