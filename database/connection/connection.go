package connection

import (
	"context"
	"log"

	"github.com/aharal/wager/configs"

	"github.com/jinzhu/gorm"
)

//ConnectionService Struct
type ConnectionService struct {
	db *gorm.DB
}

//NewDatabaseConnection Func
func NewDatabaseConnection() *ConnectionService {
	return &ConnectionService{}
}

//DBConnect Database Connection String
func (conn *ConnectionService) DBConnect() {
	db, err := gorm.Open("mysql", configs.Config.Username+":"+configs.Config.Password+"@/"+configs.Config.DatabaseName+"?charset=utf8&parseTime=True&loc=Local")
	log.Println("Connection Sucessfull!")
	if err != nil {
		configs.Ld.Logger(context.Background(), configs.ERROR, "Failed to connect database!", err)
	}
	conn.db = db
}

func (conn *ConnectionService) GetDB() *gorm.DB {
	return conn.db
}
