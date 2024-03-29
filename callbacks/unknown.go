package callbacks

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
)

func UnknownCallback(msg *tg.Message, update tg.Update, data interface{}) (tg.MessageConfig, error) {
  text := language.CurrentLanguage.Get(language.UnknownCallback)
  message := tg.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
  
  return message, nil
}

