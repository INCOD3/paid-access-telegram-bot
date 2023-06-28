package messages

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
)

func UnknownMessage(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
  text := language.CurrentLanguage.Get(language.UnknownCommand)

  replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
  replyMsg.ReplyToMessageID = update.Message.MessageID

  return replyMsg, nil
}
