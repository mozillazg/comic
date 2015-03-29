package models

import (
	"database/sql"
	"database/sql/driver"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(dbPath string) {
	if _, err := os.Stat(dbPath); err == nil {
		return
	}

	db, err := sql.Open("sqlite3", dbPath)
	checkError(err)
	defer db.Close()

	sqlStmt := `
	CREATE TABLE comic(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title CHAR(50) NOT NULL,
		date DATE NOT NULL,
		image_url CHAR(100) NOT NULL,
		description CHAR(100) NOT NULL DEFAULT ""
	);
	`
	_, err = db.Exec(sqlStmt)
	checkError(err)
}

func NewConnect(dbPath string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", dbPath)
	checkError(err)
	return
}

type Comic struct {
	ID                                 int64
	Title, ImageURL, Description, Date string
	db                                 *sql.DB
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
		values( ?, ?, ?, ?)`,
	)
	checkError(err)

	result, err := stmt.Exec(c.Title, c.ImageURL, c.Description, c.Date)
	checkError(err)
	id, err = result.LastInsertId()
	return
}

func (c *Comic) Update() (err error) {
	stmt, err := c.db.Prepare(
		`update comic set title=?, image_url=?, description=?, date=?
		where id=?`,
	)
	checkError(err)

	_, err = stmt.Exec(c.Title, c.ImageURL, c.Description, c.Date, c.ID)
	checkError(err)
	return
}

func (c *Comic) Delete() (err error) {
	_, err = c.db.Exec("delete from comic where id=?", []driver.Value{c.ID})
	checkError(err)
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
	checkError(err)

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		checkError(err)
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
		}
	}
	return
}

func LastComic(db *sql.DB) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		order by id desc limit 1`,
	)
	checkError(err)

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		checkError(err)
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
		}
	}
	return
}

func GetComic(db *sql.DB, n int64) (c *Comic, err error) {
	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		where id=? limit 1`, n,
	)
	checkError(err)

	for rows.Next() {
		var id int64
		var title string
		var imageURL string
		var description string
		var date time.Time
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		checkError(err)
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date.Format("2006-01-02"),
		}
	}
	return
}

func checkError(err error) {
	if err != nil {
		// log.Printf("%q", err)
		panic(err)
	}
}
