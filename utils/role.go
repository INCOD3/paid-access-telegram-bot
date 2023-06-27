package utils

import "github.com/w1png/paid-access-telegram-bot/language"

type Role int

const (
  Admin Role = iota
  User
)

func (r Role) String() string {
  var role string
  var err error
  switch r {
  case Admin:
    role, err = language.CurrentLanguage.Get(language.Role_Admin)
  case User:
    role, err = language.CurrentLanguage.Get(language.Role_User)
  default:
    role, err = language.CurrentLanguage.Get(language.Role_Unknown)
  }

  if err != nil {
    return err.Error()
  }
  return role
}

