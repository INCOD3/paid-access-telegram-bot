package language

import (
	"github.com/w1png/paid-access-telegram-bot/errors"
)

type Language interface {
  Get(key LanguageString) string
}

var CurrentLanguage Language

func setLanguage(language Language) {
  CurrentLanguage = language
}

func InitLanguage(language string) error {
  switch language {
  case "en":
    setLanguage(NewEnglish())
  case "ru":
    setLanguage(NewRussian())
  default:
    return errors.NewLanguageNotFoundError(language)
  }
  return nil
}

