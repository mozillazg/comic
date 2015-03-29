package main

import (
	"fmt"

	"./models"
)

func main() {
	models.CreateTable("comic.db")
	db, err := models.NewConnect("comic.db")
	db.Begin()
	defer db.Close()
	c := models.NewComic("test", "http://baidu.com", "hello", "2015-01-02", db)
	err = c.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)

	c.Title = "test2"
	c.Save()
	fmt.Println(c)

	c2 := models.NewComic("test", "http://baidu.com", "hello2", "2015-01-02", db)
	c2.Save()
	fmt.Println(c2)

	c3, _ := models.FirstComic(db)
	fmt.Println(c3)
	c4, _ := models.LastComic(db)
	fmt.Println(c4)
	c5, _ := models.GetComic(db, 1)
	fmt.Println(c5)

}
