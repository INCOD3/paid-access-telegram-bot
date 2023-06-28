package utils

import "github.com/w1png/paid-access-telegram-bot/language"

type Role int

const (
  Admin Role = iota
  User
)

func (r Role) String() string {
  var role language.LanguageString
  switch r {
  case Admin:
    role = language.Role_Admin
  case User:
    role = language.Role_User
  default:
    role = language.Role_Unknown
  }

  return language.CurrentLanguage.Get(role)
}

