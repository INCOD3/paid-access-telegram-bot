package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/states"
)

type ConsoleLogger struct {}

func NewConsoleLogger() *ConsoleLogger {
  return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(level LogLevel, message string) {
  log.Printf("[%s] %s\n", level.String(), message)
  if level == Fatal {
    os.Exit(1)
  }
}

func (l *ConsoleLogger) LogUpdate(update tg.Update, startTime time.Time) {
  username := "unknown username"
  text := "unknown text or data"
  stateText := ""

  var userId int64
  var chatId int64

  if update.Message != nil {
    username = update.Message.From.UserName
    text = fmt.Sprintf("Message text: %s", update.Message.Text)
    userId = update.Message.From.ID
    chatId = update.Message.Chat.ID
  } else if update.CallbackQuery != nil {
    username = update.CallbackQuery.From.UserName
    text = fmt.Sprintf("Callback data: %s", update.CallbackQuery.Data)
    userId = update.CallbackQuery.From.ID
    chatId = update.CallbackQuery.Message.Chat.ID
  }

  if currentState, ok := states.StateMachineInstance.States[states.NewStateUser(userId, chatId)]; ok {
    stateText = fmt.Sprintf("[State: %s]", currentState.String())
  }

  log.Printf("%s[%s] Update: [From: %s] [Data: %s] [Took: %s]\n", stateText, Info.String(), username, text, time.Since(startTime))
}

