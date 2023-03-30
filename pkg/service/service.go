package service

import (
	"github.com/Egor-Tihonov/Test-Task/pkg/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service struct {
	Key string
	DB  *db.DB
	Bot *tgbotapi.BotAPI
}

// NewService: create new service for get weather
func NewService(apikey string, db *db.DB, bot *tgbotapi.BotAPI) *Service {
	return &Service{
		Key: apikey,
		DB:  db,
		Bot: bot,
	}
}
