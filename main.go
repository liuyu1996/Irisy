package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:18692929683@tcp(127.0.0.1:3306)/ly?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connnect success")
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	stmt, errSql := db.Prepare(sql)
	if errSql != nil {
		fmt.Println(errSql)
	}
	retult, errStmt := stmt.Exec("11",1,"http://123", "http:111")
	if errStmt != nil {
		fmt.Println(errStmt)
		return
	}
	fmt.Println(retult.LastInsertId())
	return
}
