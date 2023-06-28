package language

import "github.com/w1png/paid-access-telegram-bot/errors"

type Russian struct {
  Values map[LanguageString]string
}

func NewRussian() *Russian {
  return &Russian{
    Values: map[LanguageString]string{
      Start: "Привет, это стартовое сообщение",
      Help: "Привет, это сообщение помощи",
      UnknownCommand: "Неизвестная команда",
      UnknownCallback: "Неизвестный коллбэк",
      UnknownError: "Неизвестная ошибка",
      Role_Admin: "Администратор",
      Role_User: "Пользователь",
      Role_Unknown: "Неизвестная роль",
    },
  }
}

func (e *Russian) Get(key LanguageString) string {
  value, ok := e.Values[key]
  if !ok {
    panic(errors.NewLanguageStringError(key.String()))
  }
  return value
}


