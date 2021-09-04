package service

import (
	"strings"

	"github.com/masdikaid-rakamin-homework-movie-service/contract"
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
	m.Slug = strings.ReplaceAll(m.Slug, " ", "-")
	res := databases.MysqlDB.Create(m)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *Movie) update() (*Movie, error) {
	m.Slug = strings.ReplaceAll(m.Slug, " ", "-")
	err := databases.MysqlDB.Save(m)
	if err.Error != nil {
		return nil, err.Error
	}
	return m, nil
}

func (m *Movie) UpdateFromData(newMovie Movie) (*Movie, error) {
	m.Title = newMovie.Title
	m.Slug = newMovie.Slug
	m.Description = newMovie.Description
	m.Duration = newMovie.Duration
	m.Image = newMovie.Image
	return m.update()
}

func (m *Movie) Delete() error {
	err := databases.MysqlDB.Delete(m)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (m *Movie) Contract() contract.MovieContract {
	mov := contract.MovieContract{Id: m.ID, Title: m.Title, Description: m.Description, Slug: m.Slug, Duration: m.Duration, Image: m.Image}

	return mov

}

func GetMovie(slug string) (*Movie, error) {
	var m *Movie
	err := databases.MysqlDB.First(&m, "slug=?", slug)
	if err.Error != nil {
		return nil, err.Error
	}
	return m, nil
}
