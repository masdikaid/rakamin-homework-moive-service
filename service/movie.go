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

func (m *Movie) Update() (*Movie, error) {
	err := databases.MysqlDB.Save(m)
	if err.Error != nil {
		return nil, err.Error
	}
	return m, nil
}

func (m *Movie) Delete() error {
	err := databases.MysqlDB.Delete(m)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetMovie(slug string) (*Movie, error) {
	var m *Movie
	err := databases.MysqlDB.First(&m, "slug=?", slug)
	if err.Error != nil {
		return nil, err.Error
	}
	return m, nil
}
