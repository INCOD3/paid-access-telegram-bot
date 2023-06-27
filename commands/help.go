package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

type RandomCallbackData struct {
  Name string `json:"n"`
  Surname string `json:"s"`
}

func HelpCommand(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
  text, err := language.CurrentLanguage.Get(language.Help)
  if err != nil {
    return tg.MessageConfig{}, err
  }

  replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
  replyMsg.ReplyToMessageID = update.Message.MessageID


  callback, err := utils.MarshalCallback(utils.NewCallback(
    "help",
    RandomCallbackData{
      Name: "John",
      Surname: "Doe",
    },
  ))
  if err != nil {
    return tg.MessageConfig{}, err
  }

  replyMsg.ReplyMarkup = tg.InlineKeyboardMarkup{
    InlineKeyboard: [][]tg.InlineKeyboardButton{
      []tg.InlineKeyboardButton{
        tg.InlineKeyboardButton{
          Text: "Random",
          CallbackData: &callback,
        },
      },
    },
  }

  return replyMsg, nil
}
