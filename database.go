package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Id       string `gorm:"PRIMARY_KEY"`
	Name     string
	Password string
	Status   string `gorm:"DEFAULT:1"`
}

func ConnectToDatabase() {
	db, err := gorm.Open("mysql",
		"root:root@tcp(localhost:3306)/dayu?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db)
}
