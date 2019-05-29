package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var db *gorm.DB

type ParentModel struct {
	Id        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

func ConnectToDatabase() {
	var err error
	db, err = gorm.Open("mysql",
		"root:root@tcp(localhost:3306)/dayu?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln(err)
	}

	db.SingularTable(true)
}
