package databases

import (
	"github.com/masdikaid-rakamin-homework-movie-service/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func Init(application *app.Application) {
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(application.Config.DB_Username + ":" + application.Config.DB_Password + "@tcp(" + application.Config.DB_Host + ":" + application.Config.DB_Port + ")/" + application.Config.DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("failed to connect database")
	}
}
