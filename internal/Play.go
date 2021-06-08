package internal

import (
	"MatchBot/db"
	"MatchBot/types"
	"fmt"
	"log"
	"strings"
)

var playList map[int64]types.Play = make(map[int64]types.Play)

func AddPlayer(p *types.Play, player *types.Player) {
	err := db.CreateNewUser(*player, p.ChatID)
	if err != nil {
		println(err)
		return
	}

	refreshPlayers(p)
}

func DelPlayer(p *types.Play, player *types.Player) {
	db.DeleteUser(*player, p.ChatID)
	refreshPlayers(p)
}

func refreshPlayers(p *types.Play) {
	p.Players = db.GetPlayers(p.ChatID)
}

func GetListTeam(p *types.Play) string {
	listTeam := make([]string, 0, len(p.Players))
	listTeam = append(listTeam, "Состав:")

	counter := 0
	fmt.Println(p.Players)
	for _, val := range p.Players {
		counter++
		firstName := val.FirstName
		if firstName == "" {
			firstName = val.UserName
		}
		value := fmt.Sprintf("%d %s", counter, firstName)
		listTeam = append(listTeam, value)
	}
	result := strings.Join(listTeam, "\n")
	return result
}
func CreatePlay(chatID int64) types.Play {
	log.Printf("Created new play, chatID = %d",chatID)
	return types.Play{
		ChatID:  chatID,
		Players: []types.Player{},
	}
}

func ClosePlay(chatID int64) {
	db.ClearPlayData(chatID)
	_, valExist := playList[chatID]
	if valExist {
		delete(playList, chatID)
	}
}

func GetPlay(chatID int64) types.Play {
	curentPlay, valExist := playList[chatID]
	if !valExist {
		curentPlay = CreatePlay(chatID)
		playList[chatID] = curentPlay
	}

	curentPlay.Players = append(curentPlay.Players, db.GetPlayers(chatID)...)
	return curentPlay
}
