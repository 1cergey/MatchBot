package config

import (
	"fmt"
	"github.com/joho/godotenv"

)

var TelegramToken,
DB_User,
DB_Pass,
DB_Name,
DB_Host,
DB_Port string		  

func Init() {
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error loading .env file %s", err))
	}
	fmt.Println(env)
	TelegramToken = env["TELEGRAM_TOKEN"]
	DB_User = env["DB_USER"]
	DB_Pass = env["DB_PASS"]
	DB_Name = env["DB_NAME"]
	DB_Host = env["DB_HOST"]
	DB_Port = env["DB_PORT"]
}


