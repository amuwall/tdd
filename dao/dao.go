package dao

import (
	"log"
	"tdd/config"
	"tdd/dao/client"
)

type DBManager struct {
	MySQL *client.MySQL
}

var dbManager = &DBManager{}

func Init(config config.DatabaseTemplate) (err error) {
	dbManager.MySQL, err = client.NewMySQL(config.MySQL)
	if err != nil {
		return
	}

	return
}

func Close() {
	if dbManager.MySQL != nil {
		err := dbManager.MySQL.Close()
		if err != nil {
			log.Printf("close mysql connection error %s\n", err)
		}
	}
}

func GetDB() *DBManager {
	return dbManager
}
