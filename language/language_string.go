package language

type LanguageString int

const (
  Start LanguageString = iota
  Help
  UnknownCommand
  UnknownCallback
  UnknownError
  Role_Admin
  Role_User
  Role_Unknown
  
  AdminMenu
  AdminMenu_Channels
  AdminMenu_Channels_Add
  AdminMenu_Channels_Add_Id
  AdminMenu_Channels_Add_Name
  AdminMenu_Channels_Add_Description
  AdminMenu_Channels_Add_Price
  AdminMenu_Channels_Add_IsIndefinite
  AdminMenu_Channels_Add_IsIndefinite_True
  AdminMenu_Channels_Add_IsIndefinite_False

  AdminMenu_Channels_Add_MaxMembers
  AdminMenu_Channels_Add_MaxMembers_Unlimited

  AdminMenu_Channels_Edit
  AdminMenu_Channels_Edit_ChooseChannel
  AdminMenu_Channels_Edit_Id
  AdminMenu_Channels_Edit_Name
  AdminMenu_Channels_Edit_Description
  AdminMenu_Channels_Edit_Price
  AdminMenu_Channels_Edit_IsIndefinite
  AdminMenu_Channels_Edit_MaxMembers

  AdminMenu_Channels_Remove
  AdminMenu_Channels_Remove_ChooseChannel
  AdminMenu_Channels_Remove_Confirm
)

func (l LanguageString) String() string {
  return [...]string{
    "Start",
    "Help",
    "UnknownCommand",
    "UnknownCallback",
    "UnknownError",
    "Role_Admin",
    "Role_User",
    "Role_Unknown",
  }[l]
}

