package service

import (
	"github.com/masdikaid-rakamin-homework-movie-service/databases"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `json:"title"`
	Slug        string `gorm:"unique" json:"slug"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Image       string `json:"image"`
}

func (m *Movie) Create() error {
	res := databases.MysqlDB.Create(m)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
