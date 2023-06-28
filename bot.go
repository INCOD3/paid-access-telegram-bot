package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/paid-access-telegram-bot/callbacks"
	userCallbacks "github.com/w1png/paid-access-telegram-bot/callbacks/user"
	"github.com/w1png/paid-access-telegram-bot/commands"
	"github.com/w1png/paid-access-telegram-bot/errors"
	"github.com/w1png/paid-access-telegram-bot/language"
	"github.com/w1png/paid-access-telegram-bot/logger"
	"github.com/w1png/paid-access-telegram-bot/messages"
	"github.com/w1png/paid-access-telegram-bot/models"
	"github.com/w1png/paid-access-telegram-bot/states"
	"github.com/w1png/paid-access-telegram-bot/storage"
	"github.com/w1png/paid-access-telegram-bot/utils"
)

type Bot struct {
  Bot *tg.BotAPI
  timeout int
}

func NewBot(token string, timeout int, debug bool) (*Bot, error) {
  bot, err := tg.NewBotAPI(token)
  if err != nil {
    return nil, err
  }
  bot.Debug = debug
  return &Bot{Bot: bot, timeout: timeout}, nil
}

func (b *Bot) Run() error {
  u := tg.NewUpdate(0)
  u.Timeout = b.timeout

  updates := b.Bot.GetUpdatesChan(u)
  time.Sleep(time.Millisecond * 500)
  updates.Clear()

  for update := range updates {
    go func(update tg.Update) {
      b.HandleUpdate(update)
    }(update)
  }

  return nil
}

func (b *Bot) Stop() {
  b.Bot.StopReceivingUpdates()
}

func (b *Bot) HandleUpdate(update tg.Update) {
  if update.Message != nil && (update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup()) {
    return
  }

  startTime := time.Now()

  var msg tg.MessageConfig
  var err error
  var shouldEdit bool
  var editMessage tg.Message

  // get user and create if not exists
  var userId int64
  if update.Message != nil {
    userId = update.Message.From.ID
  } else if update.CallbackQuery != nil {
    userId = update.CallbackQuery.From.ID
  }
  user, err := storage.CurrentStorage.GetUserByTelegramId(userId)
  if reflect.TypeOf(err) == reflect.TypeOf(&errors.ObjectNotFoundError{}) {
    user = models.NewUser(userId)
    err = storage.CurrentStorage.SaveUser(user)
    if err != nil {
      logger.CurrentLogger.Log(logger.Error, fmt.Sprintf("error while saving user: %s", err))
    }
  } else if err != nil {
    logger.CurrentLogger.Log(logger.Error, fmt.Sprintf("error while getting user: %s", err))
  }

  // callbacks
  if update.CallbackQuery != nil {
    callback, err := utils.NewCallbackFromCallbackData(update.CallbackQuery.Data)
    currentState, ok := states.StateMachineInstance.States[states.NewStateUser(
      user.TelegramID,
      update.CallbackQuery.Message.Chat.ID,
    )]

    if err == nil {
      if ok {
        msg, err = currentState.OnCallback(user.TelegramID, update.CallbackQuery.Message.Chat.ID, callback)
      } else {
        switch callback.Call {
        case "help":
          msg, err = userCallbacks.HelpCallback(update.CallbackQuery.Message, update, callback.Data)
          logger.CurrentLogger.Log(logger.Info, fmt.Sprintf("err: %s | msg text: %s", err, msg.Text))
          shouldEdit = true
        default:
          msg, err = callbacks.UnknownCallback(update.Message, update, callback.Data)
          shouldEdit = true
        }

        editMessage = *update.CallbackQuery.Message
      }
    }
  }

  // messages
  if update.Message != nil && !update.Message.IsCommand() {
    currentState, ok := states.StateMachineInstance.States[states.NewStateUser(
      user.TelegramID,
      update.Message.Chat.ID,
    )]

    if ok {
      msg, err = currentState.OnMessage(user.TelegramID, update.Message.Chat.ID, update.Message.Text)
    } else {
      switch update.Message.Text {
      default:
        msg, err = messages.UnknownMessage(update.Message, update)
      }
    }

    editMessage = *update.Message
  }

  // commands
  if update.Message != nil && update.Message.IsCommand() {
    switch update.Message.Command() {
    case "start":
      msg, err = commands.StartCommand(update.Message, update, user)
    case "help":
      msg, err = commands.HelpCommand(update.Message, update)
    case "test":
      msg, err = commands.TestCommand(update.Message, update)
    default:
      msg, err = commands.UnknownCommand(update.Message, update)
    }

    editMessage = *update.Message
  }

  // if no error send or edit message
  if err == nil {
    if msg.ReplyToMessageID == -1 {
      return
    }

    if shouldEdit {
      markup := tg.NewInlineKeyboardMarkup([]tg.InlineKeyboardButton{})
      if msg.ReplyMarkup != nil {
        markup = msg.ReplyMarkup.(tg.InlineKeyboardMarkup)
      }

      if _, err = b.Bot.Send(tg.NewEditMessageTextAndMarkup(
        editMessage.Chat.ID,
        editMessage.MessageID,
        msg.Text,
        markup,
      )); err != nil {
        logger.CurrentLogger.Log(logger.Error, err.Error())
      }
    } else {
      if _, err = b.Bot.Send(msg); err != nil {
        log.Println(err)
      }
    }
  }

  // if error occured during callback or command processing
  if err != nil {
    logger.CurrentLogger.Log(logger.Error, err.Error())
    text := language.CurrentLanguage.Get(language.UnknownError)
    var chatId int64
    var messageId int
    if update.Message != nil {
      chatId = update.Message.Chat.ID
      messageId = update.Message.MessageID
    } else if update.CallbackQuery != nil {
      chatId = update.CallbackQuery.Message.Chat.ID
      messageId = update.CallbackQuery.Message.MessageID
    }
    msg = tg.NewMessage(chatId, text)
    msg.ReplyToMessageID = messageId
    
    if _, err = b.Bot.Send(msg); err != nil {
      logger.CurrentLogger.Log(logger.Error, err.Error())
    }
  }

  logger.CurrentLogger.LogUpdate(update, startTime)
}

