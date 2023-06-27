package main

import (
	"log"

	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/logger"
	"github.com/w1png/paid-access-telegram-bot/states"
	"github.com/w1png/paid-access-telegram-bot/storage"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

func main() {
  config := &utils.Config{}
  err := config.GatherVariables()
  if err != nil {
    log.Fatal(err)
  }

  states.InitStateMachine()

  err = logger.InitLogger(config.LoggerType)
  if err != nil {
    log.Fatal(err)
  }

  err = storage.InitStorage(config.StorageType)

  err = language.InitLanguage(config.Language)
  if err != nil {
    logger.CurrentLogger.Log(logger.Fatal, err.Error())
  }

  bot, err := NewBot(config.TelegramToken, 60, false)
  if err != nil {
    logger.CurrentLogger.Log(logger.Fatal, err.Error())
  }

  log.Printf("Bot started as @%v\n", bot.Bot.Self.UserName)
  if err := bot.Run(); err != nil {
    logger.CurrentLogger.Log(logger.Fatal, err.Error())
  }

  log.Println("Bot stopped")
}
