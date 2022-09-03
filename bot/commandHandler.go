package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/maxwww/rd/units"
	"sort"
	"strings"
	"time"
)

const (
	commandStart       = "start"
	commandSubscribe   = "subscribe"
	commandUnsubscribe = "unsubscribe"
	commandTimetable   = "timetable"
	commandToday       = "today"
	commandTomorrow    = "tomorrow"
	commandSchedule    = "schedule"
	commandHelp        = "help"
)

func (bot *Bot) handleCommand(command string, chatId int64, user *units.User) {
	switch command {
	case commandStart, commandHelp:
		bot.handleStartCommand(chatId)
	case commandSchedule:
		bot.handleScheduleCommand(chatId)
	case commandToday:
		bot.handleTodayCommand(chatId)
	case commandTomorrow:
		bot.handleTomorrowCommand(chatId)
	case commandTimetable:
		bot.handleTimetableCommand(chatId)
	case commandSubscribe:
		bot.handleSubscribeCommand(chatId, user)
	case commandUnsubscribe:
		bot.handleUnsubscribeCommand(chatId, user)
	default:
		bot.handleUnknownCommand(chatId)
	}
	//msg := tgbotapi.NewMessage(chatId, "Сталася помилка. Спробуйте пізніше.")
	//bot.BotAPI.Send(msg)
}

func (bot *Bot) handleStartCommand(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, `Я, 🤖. Я можу допомагати тобі слідкувати за шкільним розкладом.
Я знаю розклад уроків та дзвінків у 4 класі для 21-ї школи.
Також я можу інформувати тебе про початок уроків.

Ось список моїх команд:
/today - переглянути розклад уроків на сьогодні
/tomorrow - переглянути розклад уроків на завтра
/timetable - переглянути розклад уроків на весь тиждень
/schedule - переглянути розклад дзвінків
/subscribe - увімкнути сповіщення
/unsubscribe - вимкнути сповіщення

Залишились питання чи є пропозиція? Звертайся до цього контакту - @msfilo`)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleScheduleCommand(chatId int64) {
	responseText := "<b>Розклад дзвінків:</b>\n"
	for _, bell := range Bells {
		responseText += fmt.Sprintf("%d. %s\t-\t%s\n", bell.Position, bell.Start, bell.End)
	}
	msg := tgbotapi.NewMessage(chatId, responseText)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleTodayCommand(chatId int64) {
	weekday := int(time.Now().Weekday())
	weekdayName := DayNames[weekday]
	todayTimetable, ok := Timetable[weekday]
	if !ok {
		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Ти що? Сьогодні ж %s!\nСьогодні немає уроків.", weekdayName))
		msg.ParseMode = "html"
		bot.BotAPI.Send(msg)
		return
	}
	responseText := fmt.Sprintf("<b>Розклад уроків (%s):</b>\n", weekdayName)

	responseText += getOneDayTimetable(todayTimetable.Subjects)
	msg := tgbotapi.NewMessage(chatId, responseText)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleTomorrowCommand(chatId int64) {
	weekday := int(time.Now().Weekday())
	tomorrowWeekday := weekday + 1
	if tomorrowWeekday > 6 {
		tomorrowWeekday = 0
	}
	tomorrowWeekdayName := DayNames[tomorrowWeekday]
	tomorrowTimetable, ok := Timetable[tomorrowWeekday]
	if !ok {
		mondayTimetable := Timetable[1]
		responseText := fmt.Sprintf("Завтра - %s! Вихідний!\nАле ось тобі розклад на понеділок:\n", tomorrowWeekdayName)

		responseText += getOneDayTimetable(mondayTimetable.Subjects)
		msg := tgbotapi.NewMessage(chatId, responseText)
		msg.ParseMode = "html"
		bot.BotAPI.Send(msg)
		return
	}

	responseText := fmt.Sprintf("<b>Розклад уроків (%s):</b>\n", tomorrowWeekdayName)

	responseText += getOneDayTimetable(tomorrowTimetable.Subjects)
	msg := tgbotapi.NewMessage(chatId, responseText)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleTimetableCommand(chatId int64) {
	responseText := "<b>Розклад уроків:</b>\n\n"

	keys := make([]int, 0, len(Timetable))
	for k := range Timetable {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		responseText += fmt.Sprintf("%s:\n", Timetable[k].Name)
		responseText += getOneDayTimetable(Timetable[k].Subjects)
		responseText += "\n"
	}

	msg := tgbotapi.NewMessage(chatId, responseText)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleSubscribeCommand(chatId int64, user *units.User) {
	notify := true
	err := bot.userService.UpdateUser(context.Background(), user, units.UserPatch{
		Notify: &notify,
	})
	if err != nil {
		bot.sendGeneralError(chatId)
		return
	}

	msg := tgbotapi.NewMessage(chatId, `Сповіщення увімкнено.
Тепер за 5 хв до початку уроку ти будеш отримувати сповіщення.
Аби вимкунити сповіщення скористайся /unsubscribe командою.`)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleUnsubscribeCommand(chatId int64, user *units.User) {
	notify := false
	err := bot.userService.UpdateUser(context.Background(), user, units.UserPatch{
		Notify: &notify,
	})
	if err != nil {
		bot.sendGeneralError(chatId)
		return
	}

	msg := tgbotapi.NewMessage(chatId, `Сповіщення вимкнуто.
Аби увімкнути сповіщення скористайся /subscribe командою.`)
	msg.ParseMode = "html"
	bot.BotAPI.Send(msg)
}

func (bot *Bot) handleUnknownCommand(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "На жаль, я не знаю такої команди. Скористайтеся меню або довідкою - /help")
	bot.BotAPI.Send(msg)
}

func getOneDayTimetable(subjects [][]Subject) string {
	responseText := ""

	for i, subjects := range subjects {
		var names []string
		for _, subject := range subjects {
			names = append(names, subject.ShortName)
		}

		responseText += fmt.Sprintf("%d. %s\n", i+1, strings.Join(names, "/"))
	}

	return responseText
}
