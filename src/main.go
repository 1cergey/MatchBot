package main

import (
	"fmt"
	"log"

	"github.com/1cergey/MatchBot/config"

	"github.com/joho/godotenv"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Не найдены файлы c расширением .env")
	}
}

func main() {
	cfg := config.New()
	fmt.Println(cfg.TelegramToken)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("Autorized on account", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Привет, это я бот Вася!"+"\n"+" Ты писал "+update.Message.Text,
		))
	}

}
