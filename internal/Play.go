package internal

import (
	"fmt"
	"strings"
	"MatchBot/types"
	"MatchBot/db"
	
)

type Play struct {
	ChatID int64
	Players []types.Player
}


func (p *Play) addNewPlayer(player *types.Player) {
	// db_value = types.Player{}
    // err = db.Get(&db_value, "SELECT * FROM users WHERE username=$1", player.UserName)
	// if usrExist {
	// 	return
	// } 
	
	err:=db.CreateNewUser(*player,p.ChatID)
	if err != nil   {
		println(err)
		return
	}
	
	players:= db.GetPlayers(p.ChatID)
  	// p.Players = append(p.Players, players...)
	p.Players = players  
}

func (p *Play) delPlayer(player types.Player) {
	// _, usrExist := p.Players[player.UserName]
	// if !usrExist {
	// 	return
	// }
	db.DeleteUser(player,p.ChatID) 
	//delete(p.Players,player.UserName)
	players:= db.GetPlayers(p.ChatID)
  	// p.Players = append(p.Players, players...)
	p.Players = players  

}

func (p *Play) GetListTeam() string {
	listTeam:= make([]string, 0, len(p.Players))
	listTeam = append(listTeam, "Состав:")

	counter:=0
	fmt.Println(p.Players)
	for _, val := range p.Players {
		counter++
		firstName:= val.FirstName
		if (firstName=="") {
			firstName = val.UserName
		}
		value:= fmt.Sprintf("%d %s",counter,firstName)
		listTeam = append(listTeam, value)
	}
	result := strings.Join(listTeam, "\n")
	return result
}
func CreatePlay() Play {
	var play Play = Play{}   
	return play
}

func GetPlay(chatID int64) Play {
	var CurentPlay Play = Play{
		ChatID: chatID,
		Players: []types.Player{},
	}

	players:= db.GetPlayers(chatID)
  	CurentPlay.Players = append(CurentPlay.Players, players...)
	return CurentPlay
}
