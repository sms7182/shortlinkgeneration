package database

import (
	"database/sql"
	"log"
	"shortlinkapi/models"
)

type ShortLinkDB interface {
	Open() error
	Close() error
	CreateShortLink(p *models.ShortUrlLink) error
	GetShortLinks() ([]*models.ShortUrlLink, error)
}

type DB struct {
	db *sql.DB
}

func (d *DB) Open() error {
	pg, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")
	defer pg.Close()
	err = pg.Ping()
	if err != nil {
		return err
	}

	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
