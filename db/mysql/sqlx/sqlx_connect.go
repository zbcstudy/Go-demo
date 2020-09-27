package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

const url string = "root:sneakerhead@tcp(10.0.0.131:3306)/vo_account"

type mtText struct {
	id      int    `db:"id"`
	content string `db:"content"`
	creater string `db:"creater"`
}

var DB *sqlx.DB

func main() {
	initSqlDB()
	query()
}

func initSqlDB() {
	var err error
	DB, err = sqlx.Open("mysql", url)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		log.Fatal("DB.Ping = ", err)
	}
}

func query() {
	var text = &mtText{}
	var id int
	var content string
	var creater string
	rows, _ := DB.Queryx("select id, content, creater from mt_text where id = ?", 22)
	if rows.Next() {
		rows.Scan(&id, &content, &creater)
		err := rows.StructScan(text)
		if err != nil {
			fmt.Println("struct error:", err)
		}
		fmt.Println(id, content, creater)
		fmt.Printf("%v", text)
	}
	//fmt.Printf("id:%d content:%s created:%d\n", )
}
