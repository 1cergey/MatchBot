package internal

import (
	"MatchBot/types"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//MessageHandler send text messages
func MessageHandler(u tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var replyText string

	messageText := u.Message.Text
	chatID := u.Message.Chat.ID
	usr := *u.Message.From

	if messageText == "" {
		return
	}

	msgWords := strings.Fields(messageText)
	currentPlay := GetPlay(chatID)

	switch {
	case u.Message.CommandWithAt() == "start":
		replyText = fmt.Sprintf("Привет, %s! Ну что погнали!", u.Message.From.FirstName)
	case u.Message.CommandWithAt() == "stop":
		ClosePlay(chatID)
	case len(msgWords) == 1 && strings.EqualFold(msgWords[0], "Состав"):
		replyText = GetListTeam(&currentPlay)
	case len(msgWords) == 1 && msgWords[0] == "+":
		log.Printf("player %v",usr)
		player := types.Player{UserID: strconv.Itoa(usr.ID), UserName: usr.UserName, FirstName: usr.FirstName, LastName: usr.LastName}
		AddPlayer(&currentPlay, &player)
		replyText = GetListTeam(&currentPlay)
	case len(msgWords) == 1 && msgWords[0] == "-":
		player := types.Player{UserID: strconv.Itoa(usr.ID),UserName: usr.UserName,FirstName: usr.FirstName, LastName: usr.LastName}
		DelPlayer(&currentPlay, &player)
		replyText = GetListTeam(&currentPlay)
	case len(msgWords) == 2 && msgWords[0] == "+":
		player := types.Player{UserID:msgWords[1], FirstName: msgWords[1]}
		AddPlayer(&currentPlay, &player)
		replyText = GetListTeam(&currentPlay)
	case len(msgWords) == 2 && msgWords[0] == "-":
		player := types.Player{UserID:msgWords[1],FirstName: msgWords[1]}
		DelPlayer(&currentPlay, &player)
		replyText = GetListTeam(&currentPlay)
	}

	bot.Send(tgbotapi.NewMessage(chatID, replyText))

}
