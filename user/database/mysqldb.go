package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/push?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query("select psncode, psnname, mobile, id from tb_employee where psncode = 'E009989'")
	checkErr(err)

	for rows.Next() {
		var psncode string
		var psnname string
		var mobile string
		var id int
		err = rows.Scan(&psncode, &psnname, &mobile, &id)
		fmt.Println(psncode)
		fmt.Println(psnname)
		fmt.Println(mobile)
		fmt.Println(id)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
