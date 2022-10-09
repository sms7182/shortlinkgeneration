package database

import (
	"database/sql"
	"fmt"

	"log"
	"shortlinkapi/models"

	_ "github.com/lib/pq"
)

type ShortLinkDB interface {
	Open() error
	Close() error
	CreateShortLink(p *models.ShortUrlLink) error
	GetShortLinks() ([]*models.ShortUrlLink, error)
	GetLastCounter() (counter int64, err error)
	GetActualUrl(shortLink string) (url *string, err *error)
}

type DB struct {
	db *sql.DB
}

func (d *DB) Open() error {
	pg, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}

	err = pg.Ping()
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) CreateShortLink(shortLink *models.ShortUrlLink) error {
	res, err := d.db.Exec(insertShortUrlLink, shortLink.Url, shortLink.ShortLink, shortLink.Count)
	if err != nil {
		return err
	}
	res.LastInsertId()
	return err
}

func (d *DB) GetActualUrl(shortLink string) (url string, err error) {
	var shortlink string
	qerr := d.db.QueryRow(fmt.Sprintf(getActualUrl, shortLink)).Scan(&shortlink)
	if err != nil {
		return "", qerr
	}

	return shortlink, nil
}

func (d *DB) GetLastCounter() (counter int64, err error) {
	var s sql.NullInt64
	err = d.db.QueryRow(getLastShortLink).Scan(&s)
	if err != nil {
		return 0, err
	}
	fmt.Print(s.Int64)
	return s.Int64, nil

}
