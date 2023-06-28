package states

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

type State interface {
  OnEnter(id int64, chatID int64) (tg.MessageConfig, error)
  OnExit(id int64, chatID int64) (tg.MessageConfig, error)
  OnMessage(id int64, chatID int64, message string) (tg.MessageConfig, error)
  OnCallback(id int64, chatID int64, callback utils.Callback) (tg.MessageConfig, error)

  String() string
}

