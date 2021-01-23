package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	//proxyStr := "socks5://81.17.20.50:34916"
	//proxyURL, err := url.Parse(proxyStr)

	//if err != nil {
	//	log.Println(err)
	//}

	//adding the proxy settings to the Transport object
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}

	//	client := &http.Client{
	//		Transport: transport,
	//	}
	//client.SetProxy("37.221.66.102:24531")

	client := &http.Client{}
	token := "1213192600:AAEUiACF-DSiJAmsvvj_dZ30_3iVXrTFaMo"
	bot, err := tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, client)
	//	bot, err := tgbotapi.NewBotAPI("1213192600:AAEUiACF-DSiJAmsvvj_dZ30_3iVXrTFaMo")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Autorized: %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)

	ucfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		/*if (update.Message.Text == "Stop") {
			break
		} */

		UserName := update.Message.From.UserName
		ChatID := update.Message.Chat.ID
		Text := update.Message.Text

		log.Printf("[%s] %d %s", UserName, ChatID, Text)

		reply := Text

		msg := tgbotapi.NewMessage(ChatID, reply)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
