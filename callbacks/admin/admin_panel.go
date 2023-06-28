package admincallbacks

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func AdminMenuCallback(msg *tg.Message, update tg.Update, data interface{}) (tg.MessageConfig, error) {
  message := tg.NewMessage(update.CallbackQuery.Message.Chat.ID, language.CurrentLanguage.Get(language.AdminMenu))

  callback, err := utils.MarshalCallback(
    utils.NewCallback(
      "adminmenu_channels",
      map[string]interface{}{},
    ),
  )
  if err != nil {
    return tg.MessageConfig{}, err
  }

  message.ReplyMarkup = tg.NewInlineKeyboardMarkup(
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonData(
        language.CurrentLanguage.Get(language.AdminMenu_Channels),
        callback,
      ),
    ),
  )

  return message, nil}

