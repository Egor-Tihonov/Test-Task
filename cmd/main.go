package main

import (
	"fmt"
	"runtime"
	"strings"

	config "github.com/Egor-Tihonov/Test-Task/pkg/config"
	"github.com/Egor-Tihonov/Test-Task/pkg/db"
	"github.com/Egor-Tihonov/Test-Task/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	InitLog()

	c, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("get configs, %e", err)
	}

	d, err := db.New(c.DBURL)
	if err != nil {
		logrus.Fatalf("cannot connect to db, %e", err.Error())
	}

	defer d.Pool.Close()

	bot, err := tgbotapi.NewBotAPI(c.TgApi)
	if err != nil {
		logrus.Fatalf("cannot create bot, %e", err)
	}
	logrus.Info("Start bot")
	s := service.NewService(c.ApiKey, d, bot)

	s.BotLogic()

}

func InitLog() {
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	})

}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
