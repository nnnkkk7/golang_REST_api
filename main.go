package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

type Tweet struct {
	gorm.Model
	Content string
}

func Init() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&Tweet{})
}

func dbInsert(content string) error {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(&Tweet{Content: content})
	return nil
}

func GetAll() ([]Tweet, error) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var tweets []Tweet
	db.Order("created_at desc").Find(&tweets)
	return tweets, nil

}

func main() {
	e := echo.New()

	Init()

	e.GET("/", func(c echo.Context) {
		tweets, _ := GetAll()
		s := strings.Join(tweets, ",")
		c.String(200, s)
	})

}
