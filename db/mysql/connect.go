package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type MtRole struct {
	Id        int    `db:"id"`
	companyId int    `db:"company_id"`
	name      string `db:"name"`
}

const url string = "root:sneakerhead@tcp(10.0.0.131:3306)/vo_account"
const selectRole string = "select `id`,`company_id`,`name` from mt_role"

// 初始化数据库连接
func initDB() *sql.DB {
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Print("数据库连接出错")
		panic(err)
	} else {
		fmt.Println("数据库连接success")
	}
	return db
}

func main() {
	search()
	insert()
}

func search() {
	db := initDB()
	defer db.Close()
	rows, _ := db.Query(selectRole)
	var id int
	var companyId int
	var name string
	for rows.Next() {
		rows.Columns()
		rows.Scan(&id, &companyId, &name)
		fmt.Println(id, "--", companyId, "--", name)
	}

}

func insert() {
	db := initDB()
	defer db.Close()
	stmt, err := db.Prepare("insert into mt_role(company_id,name,active,default_role,creater,created,modifier," +
		"modified) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Panicln(err)
		return
	}
	res, error := stmt.Exec(1, "test1", true, true, "zbc", time.Now(), "zbc", time.Now())
	if error != nil {
		log.Print(error)
	}
	//获取最后插入的id
	id, _ := res.LastInsertId()
	log.Println("插入获取的id为 ", id)
	affected, _ := res.RowsAffected()
	log.Println("受影响的行数 ", affected)
}
