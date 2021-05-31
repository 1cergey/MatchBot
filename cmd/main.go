package main

import (
	"MatchBot/db"
	"MatchBot/internal"
	"fmt"
	"log"

	cfg "MatchBot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Не найдены файлы c расширением .env")
	}
	cfg.Init()
	db.Connect()

}

func main() {

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("Autorized on account", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println("Ошибка при получении обновлений")
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		internal.MessageHandler(update, bot)
	}

}
