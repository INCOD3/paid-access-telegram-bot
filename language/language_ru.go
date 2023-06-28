package language

import "github.com/w1png/paid-access-telegram-bot/errors"

type Russian struct {
  Values map[LanguageString]string
}

func NewRussian() *Russian {
  return &Russian{
    Values: map[LanguageString]string{
      Start: "Привет, это стартовое сообщение",
      Help: "Привет, это сообщение помощи",
      UnknownCommand: "Неизвестная команда",
      UnknownCallback: "Неизвестный коллбэк",
      UnknownError: "Неизвестная ошибка",
      Role_Admin: "Администратор",
      Role_User: "Пользователь",
      Role_Unknown: "Неизвестная роль",

      AdminMenu: "Меню администратора",
      AdminMenu_Channels: "Каналы",
      AdminMenu_Channels_Add: "Добавить канал",
      AdminMenu_Channels_Add_Id: "Введите ID канала (например, @channel_name)",
      AdminMenu_Channels_Add_Name: "Введите название канала",
      AdminMenu_Channels_Add_Description: "Введите описание канала",
      AdminMenu_Channels_Add_Price: "Введите цену подписки на канал",
      AdminMenu_Channels_Add_IsIndefinite: "Бессрочная подписка?",
      AdminMenu_Channels_Add_IsIndefinite_True: "Да",
      AdminMenu_Channels_Add_IsIndefinite_False: "Нет",
      AdminMenu_Channels_Add_MaxMembers: "Введите максимальное количество подписчиков",
      AdminMenu_Channels_Add_MaxMembers_Unlimited: "Безлимит",

      AdminMenu_Channels_Edit: "Редактировать канал",
      AdminMenu_Channels_Edit_ChooseChannel: "Выберите канал",
      AdminMenu_Channels_Edit_Id: "Введите ID канала (например, @channel_name)",
      AdminMenu_Channels_Edit_Name: "Введите название канала",
      AdminMenu_Channels_Edit_Description: "Введите описание канала",
      AdminMenu_Channels_Edit_Price: "Введите цену подписки на канал",
      AdminMenu_Channels_Edit_IsIndefinite: "Бессрочная подписка?",
      AdminMenu_Channels_Edit_MaxMembers: "Введите максимальное количество подписчиков",

      AdminMenu_Channels_Remove: "Удалить канал",
      AdminMenu_Channels_Remove_ChooseChannel: "Выберите канал",
      AdminMenu_Channels_Remove_Confirm: "Вы уверены, что хотите удалить канал?",
    },
  }
}

func (e *Russian) Get(key LanguageString) string {
  value, ok := e.Values[key]
  if !ok {
    panic(errors.NewLanguageStringError(key.String()))
  }
  return value
}


