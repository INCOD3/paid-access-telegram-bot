package language

import "github.com/w1png/paid-access-telegram-bot/errors"

type English struct {
  Values map[LanguageString]string
}

func NewEnglish() *English {
  return &English{
    Values: map[LanguageString]string{
      Start: "Hello, this is a start message",
      Help: "Hello, this is a help message",
      UnknownCommand: "Unknown command",
      UnknownCallback: "Unknown callback",
      UnknownError: "Unknown error",
      Role_Admin: "Administrator",
      Role_User: "User",
      Role_Unknown: "Unknown role",
    },
  }
}

func (e *English) Get(key LanguageString) (string, error) {
  value, ok := e.Values[key]
  if !ok {
    return "", errors.NewLanguageStringError(key.String())
  }
  return value, nil
}

