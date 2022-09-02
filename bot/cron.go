package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/maxwww/rd/units"
	"github.com/robfig/cron/v3"
	"log"
	"strconv"
	"strings"
	"time"
)

func (bot *Bot) RegisterCrons() {
	loc, _ := time.LoadLocation("Europe/Kiev")
	c := cron.New(cron.WithLocation(loc))

	for day, tTable := range Timetable {
		for number, subject := range tTable.Subjects {
			parts := strings.Split(Bells[number].Start, ":")
			hours, _ := strconv.Atoi(parts[0])
			minutes, _ := strconv.Atoi(parts[1])

			var subjectNames []string
			for _, s := range subject {
				subjectNames = append(subjectNames, s.ShortName)
			}
			message := fmt.Sprintf(`Урок "%s" розпочався.`, strings.Join(subjectNames, "/"))
			_, err := c.AddFunc(fmt.Sprintf("%d %d * * %d", minutes, hours, day), bot.buildCronFunction(message))

			if err != nil {
				panic(err)
			}

			minutes -= 5
			if minutes < 0 {
				minutes = 60 + minutes
				hours -= 1
			}

			message = fmt.Sprintf(`Урок "%s" розпочнеться через 5 хв.`, strings.Join(subjectNames, "/"))
			_, err = c.AddFunc(fmt.Sprintf("%d %d * * %d", minutes, hours, day), bot.buildCronFunction(message))

			if err != nil {
				panic(err)
			}
		}
	}

	c.Start()
}

func (bot *Bot) buildCronFunction(message string) func() {
	return func() {
		notify := true
		users, err := bot.userService.Users(context.Background(), units.UserFilter{
			Notify: &notify,
		})
		if err != nil {
			log.Println(err)
		}
		for _, user := range users {
			msg := tgbotapi.NewMessage(int64(user.TelegramID), message)
			msg.ParseMode = "html"
			bot.BotAPI.Send(msg)
		}
	}
}
