package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewConnect(dbPath string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", dbPath)
	checkError(err)
	return
}

type Comic struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ImageURL    string `json:"imageURL"`
	Description string `json:"description"`
	Date        string `json:"date"`
	db          *sql.DB
}

func (c *Comic) Save() (err error) {
	if c.ID == 0 {
		id, err := c.Create()
		c.ID = id
		return err
	}
	return c.Update()
}

func (c *Comic) Create() (id int64, err error) {
	stmt, err := c.db.Prepare(
		`insert into comic(title, image_url, description, date)
		values( $1, $2, $3, $4)`,
	)
	if checkError(err) {
		return 0, err
	}

	result, err := stmt.Exec(c.Title, c.ImageURL, c.Description, c.Date)
	if checkError(err) {
		return 0, err
	}
	id, err = result.LastInsertId()
	return
}

func (c *Comic) Update() (err error) {
	stmt, err := c.db.Prepare(
		`update comic set title=$1, image_url=$2, description=$3, date=$4
		where id=$5`,
	)
	if checkError(err) {
		return
	}

	_, err = stmt.Exec(c.Title, c.ImageURL, c.Description, c.Date, c.ID)
	checkError(err)
	return
}

func (c *Comic) Delete() (err error) {
	_, err = c.db.Exec("delete from comic where id=$1", c.ID)
	if checkError(err) {
		return
	}
	c.ID = 0
	return
}

func NewComic(title, url, desc, date string, db *sql.DB) (c *Comic) {
	return &Comic{
		ID:          0,
		Title:       title,
		ImageURL:    url,
		Description: desc,
		Date:        date,
		db:          db,
	}
}

func FirstComic(db *sql.DB) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		order by id limit 1`,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		}
		return
	}
	return
}

func LastComic(db *sql.DB) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		order by id desc limit 1`,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		}
		return
	}
	return
}

func GetComic(db *sql.DB, n int64) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		where id=$1 limit 1`, n,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		}
		return
	}
	return
}

func RandomComic(db *sql.DB) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		order by random() limit 1`,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		}
		return
	}
	return
}

func checkError(err error) bool {
	if err != nil {
		log.Printf("%q", err)
		return true
	}
	return false
}

func AllComic(db *sql.DB) (c []*Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		order by date desc`,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = append(c, &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		})
	}
	return
}

func PrevComic(db *sql.DB, n int64) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		where id<$1 order by id desc limit 1`, n,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		}
		return
	}
	return
}

func NextComic(db *sql.DB, n int64) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		where id>$1 order by id asc limit 1`, n,
	)
	if checkError(err) {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		if checkError(err) {
			return nil, err
		}
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
			db:          db,
		}
		return
	}
	return
}
