package internal

import (
	"fmt"
	"strings"
	"MatchBot/types"
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
	currentPlay := GetPlay(chatID);

	switch {
	case u.Message.CommandWithAt() == "start":
		replyText = fmt.Sprintf("Привет, %s! Помогу Вам собрать команду мечты!", u.Message.From.FirstName)
	case len(msgWords) == 1 && strings.EqualFold(msgWords[0],"Состав"):
		replyText = currentPlay.GetListTeam()	
	case len(msgWords) == 1 && msgWords[0] == "+":
		player := types.Player{UserName:usr.UserName, FirstName:usr.FirstName, LastName: usr.LastName}
	 	currentPlay.addNewPlayer(&player)
		replyText= currentPlay.GetListTeam()	
	case len(msgWords) == 2 && msgWords[0] == "+":
		player := types.Player{UserName: msgWords[1], FirstName:msgWords[1]}
		currentPlay.addNewPlayer(&player)
		replyText = currentPlay.GetListTeam()	
	case len(msgWords) == 1 && msgWords[0] == "-":
		player := types.Player{UserName: usr.UserName}
		currentPlay.delPlayer(player)
		replyText = currentPlay.GetListTeam()	
	case len(msgWords) == 2 && msgWords[0] == "-":
		player := types.Player{UserName: msgWords[1]}
		currentPlay.delPlayer(player)
		replyText = currentPlay.GetListTeam()	
	default:
		fmt.Println(messageText)
	}

	bot.Send(tgbotapi.NewMessage(chatID, replyText))

}


