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

