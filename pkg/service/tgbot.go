package service

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/Egor-Tihonov/Test-Task/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (s *Service) BotLogic() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.Bot.GetUpdatesChan(u)

	first_req := ""

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if first_req == "" {
			dt, err := s.DB.Get(context.Background(), update.SentFrom().UserName)
			if err != nil {
				if err == models.ErrorUserDontExist {
					err := s.DB.CreateUser(context.Background(), update.SentFrom().UserName)
					if err != nil {
						logrus.Errorf("error bot logic, %e", err)
					}
				}
			}
			if dt != nil {
				first_req = dt.Format("02 January 2006")
			}
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я бот который поможет тебе узнать погоду на сегодня!\n\nДля того что бы узнать погоду в городе\nнеобходимо просто отправить боту название города\nТак же у бота есть функция вывода даты вашего первого запроса к нему /when_first_request")
				s.Bot.Send(msg)

			case "/when_first_request":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Первый запрос был \n%s", first_req))
				s.Bot.Send(msg)
				
			default:
				ms := update.Message.Text
				m, err := s.GetWeather(ms)
				if err != nil {
					logrus.Errorf("error bot logic, %e", err)
				}

				message := ""
				if m.FeelsLike < 5 {
					message = "Оденьтесь потеплее"
					if m.Status == "дождь" || m.Status == "пасмурно" {
						message += " и зонт бы тоже не помешал!"
					}
				} else if m.FeelsLike > 5 && m.FeelsLike < 18 {
					message = "На улице уже тепло, но и не жара, не раздевайтесь слишком сильно"
				} else if m.FeelsLike > 18 {
					message = "Не забудьте намазаться кремом от загара и одеть головной убор, на улице жарко"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n\nСейчас в городе %s %s, %o°C,\nТемпература ощущается как:  %o°C,\nМаксималная температура сегодня:  %o°C,\nМинимальная температура сегодня:  %o°C,\nВетер:  %.2fм/с",
					message,
					m.Country,
					m.Status,
					m.FactTemp,
					m.FeelsLike,
					m.MaxTemp,
					m.MinTemp,
					m.WindSpeed,
				))
				s.Bot.Send(msg)

				time_request := time.Now()
				if first_req == "" {
					err := s.DB.UpdateTime(context.Background(), time_request)
					if err != nil {
						logrus.Errorf("error bot logic, %e", err)
					}
					first_req = time_request.Format("Jan 2 15:04:05 2006")
				}
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
			s.Bot.Send(msg)
		}
	}
}
