package adminmessages

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func AdminMenuMessage(msg *tg.Message, update tg.Update) (tg.MessageConfig, error) {
  replyMsg := tg.NewMessage(update.Message.Chat.ID, language.CurrentLanguage.Get(language.AdminMenu))

  callback, err := utils.MarshalCallback(
    utils.NewCallback(
      "adminmenu_channels",
      map[string]interface{}{},
    ),
  )
  if err != nil {
    return tg.MessageConfig{}, err
  }

  replyMsg.ReplyMarkup = tg.NewInlineKeyboardMarkup(
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonData(
        language.CurrentLanguage.Get(language.AdminMenu_Channels),
        callback,
      ),
    ),
  )

  return replyMsg, nil
}
