package db

import (
	cfg "MatchBot/config"
	"MatchBot/types"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var db *sqlx.DB

func Connect() {
	port, _ := strconv.ParseUint(cfg.DB_Port, 10, 16)
	connConfig := pgx.ConnConfig{
		Host:                 cfg.DB_Host,
		Port:                 uint16(port),
		Database:             cfg.DB_Name,
		User:                 cfg.DB_User,
		Password:             cfg.DB_Pass,
		PreferSimpleProtocol: false,
	}
	fmt.Println(connConfig)
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 100,
		AcquireTimeout: 600 * time.Second,
		AfterConnect:   nil,
	})

	if err != nil {
		fmt.Printf("Unable to connect to database: %s", err)
	}

	_db := stdlib.OpenDBFromPool(connPool)
	_db.SetMaxIdleConns(0)
	_db.SetMaxOpenConns(100)
	db = sqlx.NewDb(_db, "pgx")
}

func CreateNewUser(user types.Player, chatId int64) error {
	log.Printf("Start creating new user, chatID = %d\n %v\n", chatId, user)

	tx, err := db.Begin()

	if err != nil {
		return errors.Wrapf(err, "error when create tx in CreateNewUser")
	}

	if UserExist(user, chatId) {
		return nil
	}

	_, err = tx.Exec(
		"INSERT INTO users (chat_id, user_id, first_name, last_name, username) VALUES ($1, $2, $3, $4, $5)",
		chatId, user.UserID,user.FirstName, user.LastName, user.UserName)
	if err != nil {
		trErr := tx.Rollback()
		if trErr != nil {
			return errors.Wrapf(trErr, "tx error when insert user first_name: %s, username: %s", user.FirstName, user.UserName)
		}
		return errors.Wrapf(err, "error when insert user first_name: %s, username: %s", user.FirstName, user.UserName)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, " error when commit tx on create user")
	}
	log.Printf("Created user %v\n",user)

	return nil
}

func DeleteUser(user types.Player, chatId int64) error {
	if !UserExist(user, chatId) {
		return nil
	}

	tx, err := db.Begin()

	if err != nil {
		return errors.Wrapf(err, "error when create tx in CreateNewUser")
	}

	_, err = tx.Exec(
		"DELETE FROM users WHERE chat_id = $1 and user_id = $2", chatId, user.UserID)
	if err != nil {
		trErr := tx.Rollback()
		if trErr != nil {
			return errors.Wrapf(trErr, "tx error when delete user first_name: %s, username: %s", user.FirstName, user.UserName)
		}
		return errors.Wrapf(err, "error when delete user first_name: %s, username: %s", user.FirstName, user.UserName)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, " error when delete tx on create user")
	}

	return nil
}

func GetPlayers(chatID int64) []types.Player {
	players := []types.Player{}
	db.Select(&players, "SELECT username, first_name, last_name  FROM users WHERE chat_id = $1", chatID)
	log.Printf("Got players, chatID = %d\n players %v\n",chatID,players)

	return players
}

func ClearPlayData(chatId int64) error {
	tx, err := db.Begin()

	if err != nil {
		return errors.Wrapf(err, "error when create tx in CreateNewUser")
	}

	_, err = tx.Exec(
		"DELETE FROM users WHERE chat_id = $1", chatId)
	if err != nil {
		tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, " error when delete tx on create user")
	}

	return nil
}

func UserExist(user types.Player, chatId int64) bool {
	err:= db.Get(&user.UserID, "SELECT user_id FROM users WHERE chat_id=$1 and user_id = $2;", chatId,user.UserID)
	if err!=nil {
		log.Printf("UserExist query error %v\n",err)
	}
	userExist:= err==nil
	log.Printf("userExist %b",userExist)
	return userExist
}
