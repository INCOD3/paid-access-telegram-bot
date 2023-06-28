package language

import "github.com/w1png/paid-access-telegram-bot/errors"

type English struct {
  Values map[LanguageString]string
}

func NewEnglish() *English {
  return &English{
    Values: map[LanguageString]string{
      Start: "Hello, this is a start message",
      Help: "Hello, this is a help message",
      UnknownCommand: "Unknown command",
      UnknownCallback: "Unknown callback",
      UnknownError: "Unknown error",
      Role_Admin: "Administrator",
      Role_User: "User",
      Role_Unknown: "Unknown role",

      AdminMenu: "Admin menu",
      AdminMenu_Channels: "Channels",
      AdminMenu_Channels_Add: "Add channel",
      AdminMenu_Channels_Add_Id: "Enter channel id in format @channel_name",
      AdminMenu_Channels_Add_Name: "Enter channel name",
      AdminMenu_Channels_Add_Description: "Enter channel description",
      AdminMenu_Channels_Add_Price: "Enter subscription price",
      AdminMenu_Channels_Add_IsIndefinite: "Is subscription indefinite?",
      AdminMenu_Channels_Add_IsIndefinite_True: "Yes",
      AdminMenu_Channels_Add_IsIndefinite_False: "No",

      AdminMenu_Channels_Add_MaxMembers: "Enter maximum number of members",
      AdminMenu_Channels_Add_MaxMembers_Unlimited: "Unlimited",

      AdminMenu_Channels_Edit: "Edit channel",
      AdminMenu_Channels_Edit_ChooseChannel: "Choose channel to edit",
      AdminMenu_Channels_Edit_Id: "Enter channel id in format @channel_name",
      AdminMenu_Channels_Edit_Name: "Enter channel name",
      AdminMenu_Channels_Edit_Description: "Enter channel description",
      AdminMenu_Channels_Edit_Price: "Enter subscription price",
      AdminMenu_Channels_Edit_IsIndefinite: "Is subscription indefinite?",
      AdminMenu_Channels_Edit_MaxMembers: "Enter maximum number of members",

      AdminMenu_Channels_Remove: "Remove channel",
      AdminMenu_Channels_Remove_ChooseChannel: "Choose channel to remove",
      AdminMenu_Channels_Remove_Confirm: "Are you sure you want to remove channel %s?",
    },
  }
}

func (e *English) Get(key LanguageString) string {
  value, ok := e.Values[key]
  if !ok {
    panic(errors.NewLanguageStringError(key.String()))
  }
  return value
}

