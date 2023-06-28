package admincallbacks

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func AdminMenuChannelCallback(msg *tg.Message, update tg.Update, data interface{}) (tg.MessageConfig, error) {
  text := language.CurrentLanguage.Get(language.AdminMenu_Channels)
  message := tg.NewMessage(
    update.CallbackQuery.Message.Chat.ID,
    text,
  )

  add_callback, err := utils.MarshalCallback(
    utils.NewCallback(
      "adminmenu_channels_add",
      map[string]interface{}{},
    ),
  )
  if err != nil {
    return tg.MessageConfig{}, err
  }

  remove_callback, err := utils.MarshalCallback(
    utils.NewCallback(
      "adminmenu_channels_remove",
      map[string]interface{}{},
    ),
  )
  if err != nil {
    return tg.MessageConfig{}, err
  }

  edit_callback, err := utils.MarshalCallback(
    utils.NewCallback(
      "adminmenu_channels_edit",
      map[string]interface{}{},
    ),
  )
  if err != nil {
    return tg.MessageConfig{}, err
  }

  back_callback, err := utils.MarshalCallback(
    utils.NewCallback(
      "adminmenu",
      map[string]interface{}{},
    ),
  )
  if err != nil {
    return tg.MessageConfig{}, err
  }

  message.ReplyMarkup = tg.NewInlineKeyboardMarkup(
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonData(
        language.CurrentLanguage.Get(language.AdminMenu_Channels_Add),
        add_callback,
      ),
    ),
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonData(
        language.CurrentLanguage.Get(language.AdminMenu_Channels_Remove),
        remove_callback,
      ),
    ),
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonData(
        language.CurrentLanguage.Get(language.AdminMenu_Channels_Edit),
        edit_callback,
      ),
    ),
    tg.NewInlineKeyboardRow(
      tg.NewInlineKeyboardButtonData(
        language.CurrentLanguage.Get(language.Back),
        back_callback,
      ),
    ),
  )

  return message, nil
}

