package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mtText struct {
	Id      int    `db:"id"`
	Content string `db:"content"`
	Creater string `db:"creater"`
	Str     string `db:"str"`
}

func (mt mtText) TableName() string {
	return "mt_text"
}

const url string = "root:sneakerhead@tcp(10.0.0.131:3306)/vo_account"

func main() {
	mt := mtText{}
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        url,
	}), &gorm.Config{})
	//_ = db.AutoMigrate(&mt)
	err := db.Migrator().AddColumn(&mt, "str")
	if err != nil {
		fmt.Println("add column fail", err)
	}
	db.First(&mt)
	fmt.Println(mt)
}
