package userCallbacks

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HelpCallback(msg *tg.Message, update tg.Update, data interface{}) (tg.MessageConfig, error) {
  name, ok := data.(map[string]interface{})["n"].(string)
  surname, ok := data.(map[string]interface{})["s"].(string)
  if !ok {
    return tg.MessageConfig{}, fmt.Errorf("Invalid data")
  }
  text := fmt.Sprintf("Hello %s %s, I am happy to help you!", name, surname)

  message := tg.NewMessage(msg.Chat.ID, text)
  return message, nil
}

