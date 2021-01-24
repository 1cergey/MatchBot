package main

import (
	"log"
	//"net/http"
	"fmt"

	"github.com/1cergey/MatchBot/config"

	"github.com/joho/godotenv"
	//tgbotapi "gopkg.in/telegram-bot-api.v5"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Не найдены файлы c расширением .env")
	}
}

func main() {
	cfg := config.New()
	fmt.Print(cfg.TelegramToken)
	return

	// client := &http.Client{}
	// token :=
	// bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, client)
	// //	bot, err := tgbotapi.NewBotAPI("1213192600:AAEUiACF-DSiJAmsvvj_dZ30_3iVXrTFaMo")

	// if err != nil {
	// 	log.Panic(err)
	// }

	// bot.Debug = true
	// log.Printf("Autorized: %s", bot.Self.UserName)

	// var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)

	// ucfg.Timeout = 60

	// updates, err := bot.GetUpdatesChan(ucfg)

	// for update := range updates {
	// 	if update.Message == nil { // ignore any non-Message Updates
	// 		continue
	// 	}

	// 	/*if (update.Message.Text == "Stop") {
	// 		break
	// 	} */

	// 	UserName := update.Message.From.UserName
	// 	ChatID := update.Message.Chat.ID
	// 	Text := update.Message.Text

	// 	log.Printf("[%s] %d %s", UserName, ChatID, Text)

	// 	reply := Text

	// 	msg := tgbotapi.NewMessage(ChatID, reply)
	// 	msg.ReplyToMessageID = update.Message.MessageID

	// 	bot.Send(msg)
	// }
}
