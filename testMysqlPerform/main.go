package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	sTime := time.Now()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "admin123", "10.23.14.50", 5000, "test")
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetMaxOpenConns(1)

	sql := "select * from vip_agent_recharge_detail_fzs"
	rows, err := conn.Query(sql)
	c, _ := rows.ColumnTypes()
	c[0].Name()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var rowNum int
	for rows.Next() {
		rowNum++
	}
	fmt.Println(time.Since(sTime))

	fmt.Println(rowNum)

}
