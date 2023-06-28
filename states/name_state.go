package states

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

type NameState struct {
  Name string
}

func (s *NameState) OnEnter(id int64, chatID int64) (tg.MessageConfig, error) {
  msg := tg.NewMessage(chatID, "What is your name?")
  return msg, nil
}

func (s *NameState) OnExit(id int64, chatID int64) (tg.MessageConfig, error) {
  msg := tg.NewMessage(chatID, "Goodbye, " + s.Name + "!")
  msg.ReplyToMessageID = -1

  StateMachineInstance.RemoveState(NewStateUser(id, chatID))
  return msg, nil
}

func (s *NameState) OnMessage(id int64, chatID int64, message string) (tg.MessageConfig, error) {
  s.Name = message
  msg := tg.NewMessage(chatID, "Nice to meet you, " + s.Name + "!")
  s.OnExit(id, chatID)
  return msg, nil
}

func (s *NameState) OnCallback(id int64, chatID int64, callback utils.Callback) (tg.MessageConfig, error) {
  return tg.MessageConfig{}, nil
}

func (s NameState) String() string {
  return "NameState"
}

func NewNameState() State {
  return &NameState{}
}
