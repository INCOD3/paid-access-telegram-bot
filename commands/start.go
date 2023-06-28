package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/models"
)

func StartCommand(msg *tg.Message, update tg.Update, user *models.User) (tg.MessageConfig, error) {
  text := language.CurrentLanguage.Get(language.Start)

  replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
  replyMsg.ReplyToMessageID = update.Message.MessageID

  if user.IsAdmin() {
    replyMsg.ReplyMarkup = tg.NewReplyKeyboard(
      tg.NewKeyboardButtonRow(
        tg.NewKeyboardButton(language.CurrentLanguage.Get(language.AdminMenu)),
      ),
    )
  }

  return replyMsg, nil
}
