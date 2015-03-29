package main

import (
	"database/sql"
	"database/sql/driver"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	checkError(err)
	defer db.Close()

	sqlStmt := `
	CREATE TABLE comic(
		id INT PRIMARY KEY AUTOINCREMENT,
		title CHAR(50) NOT NULL,
		date DATE NOT NULL,
		image_url CHAR(100) NOT NULL,
		description CHAR(100) NOT NULL DEFAULT "",
	);
	`
	_, err = db.Exec(sqlStmt)
	checkError(err)
}

func newConnect(dbPath string) (db driver.Conn, err error) {
	db, err = sql.Open("sqlite3", dbPath)
	checkError(err)
	return
}

type Comic struct {
	ID                                 int
	Title, ImageURL, Description, Date string
	db                                 driver.Conn
}

func (c *Comic) save() (err error) {
	if c.ID == 0 {
		return c.create()
	}
	return c.update()
}

func (c *Comic) create() (err error) {
	c.db.Begin()
	defer c.db.Close()
	stmt, err := db.Prepare(
		`insert into comic(title, image_url, description, data)
		values( ?, ?, ?, ?)`,
	)
	checkError(err)

	_, err = stmt.Exec(c.Title, c.ImageURL, c.Description, c.Date)
	checkError(err)
	return
}

func (c *Comic) update() (err error) {
	c.db.Begin()
	defer c.db.Close()
	stmt, err := db.Prepare(
		`update comic set title=?, image_url=?, description=?, data=?
		where id=?`,
	)
	checkError(err)

	_, err = stmt.Exec(c.Title, c.ImageURL, c.Description, c.Date, c.ID)
	checkError(err)
	return
}

func (c *Comic) delete() (err error) {
	c.db.Begin()
	defer c.db.Close()
	_, err = db.Exec("delete from comic where id=?", c.ID)
	checkError(err)
	return
}

func newComic(title, url, desc, date string, db driver.Conn) (c *Comic) {
	return &Comic{
		0, title, date, url, des, db,
	}
}

func firstComic(db driver.Conn) (c *Comic, err error) {
	db.Begin()
	defer db.Close()

	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		limit 1 order by id`,
	)
	checkError(err)

	for rows.Next() {
		var id int
		var title string
		var imageURL string
		var description string
		var date string
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		checkError(err)
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date,
		}
	}
	return
}

func lastComic(db driver.Conn) (c *Comic, err error) {
	db.Begin()
	defer db.Close()

	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		limit 1 order by id desc`,
	)
	checkError(err)

	for rows.Next() {
		var id int
		var title string
		var imageURL string
		var description string
		var date string
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		checkError(err)
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date,
		}
	}
	return
}

func getComic(db driver.Conn, n int) (c *Comic, err error) {
	db.Begin()
	defer db.Close()

	rows, err := db.Query(
		`select id, title, image_url, description, date from comic
		limit 1 where id=?`, n,
	)
	checkError(err)

	for rows.Next() {
		var id int
		var title string
		var imageURL string
		var description string
		var date string
		err = rows.Scan(&id, &title, &imageURL, &description, &date)
		checkError(err)
		c = &Comic{
			ID:          id,
			Title:       title,
			ImageURL:    imageURL,
			Description: description,
			Date:        date,
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
