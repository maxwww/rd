package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/maxwww/rd/postgres"
	"github.com/maxwww/rd/units"
	"log"
)

type Bot struct {
	BotAPI      *tgbotapi.BotAPI
	userService units.UserService
}

func NewBot(botAPI *tgbotapi.BotAPI, db *postgres.DB) *Bot {
	bot := Bot{
		BotAPI: botAPI,
	}

	bot.userService = postgres.NewUserService(db)

	return &bot
}

func (bot *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.BotAPI.GetUpdatesChan(u)

	bot.RegisterCrons()

	for update := range updates {
		go bot.handleUpdate(update)
	}

	return nil
}

func (bot *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil && update.CallbackQuery == nil { // ignore any non-Message Updates
		return
	}

	var fromUser *tgbotapi.User
	var chatId int64

	if update.CallbackQuery != nil {
		fromUser = update.CallbackQuery.From
		chatId = update.CallbackQuery.Message.Chat.ID
	} else {
		fromUser = update.Message.From
		chatId = update.Message.Chat.ID
	}

	user, err := bot.userService.UserByTelegramID(context.Background(), uint(fromUser.ID))
	if err != nil {
		if err != units.ErrNotFound {
			// TODO: handle error
			log.Println(err)
			bot.sendGeneralError(chatId)
			return
		}
		err = bot.userService.CreateUser(context.Background(), &units.User{
			TelegramID:   uint(update.Message.From.ID),
			IsBot:        update.Message.From.IsBot,
			FirstName:    update.Message.From.FirstName,
			LastName:     update.Message.From.LastName,
			UserName:     update.Message.From.UserName,
			LanguageCode: update.Message.From.LanguageCode,
			Notify:       false,
		})
		if err != nil {
			bot.sendGeneralError(chatId)
			return
		}

		user, err = bot.userService.UserByTelegramID(context.Background(), uint(fromUser.ID))

	}

	if update.Message.IsCommand() {
		bot.handleCommand(update.Message.Command(), chatId, user)
		return
	}

	bot.handleUnknownCommand(chatId)
}

func (bot *Bot) sendGeneralError(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "Сталася помилка. Спробуйте пізніше.")
	bot.BotAPI.Send(msg)
}
