package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MtRole struct {
	Id        int    `db:"id"`
	companyId int    `db:"company_id"`
	name      string `db:"name"`
}

const url string = "root:sneakerhead@tcp(10.0.0.131:3306)/vo_account"

// 初始化数据库连接
func initDB() *sql.DB {
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Print("数据库连接出错")
		panic(err)
	} else {
		fmt.Println("success")
	}
	return db
}

func main() {
	search()
}

func search() {
	db := initDB()
	defer db.Close()
	rows, _ := db.Query("select id,company_id,name from mt_role")
	var id int
	var companyId int
	var name string
	for rows.Next() {
		rows.Columns()
		rows.Scan(&id, &companyId, &name)
		fmt.Println(id, "--", companyId, "--", name)
	}
}
