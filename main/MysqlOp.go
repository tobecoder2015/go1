package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func main()  {
	db,err:=sql.Open("mysql","user:password@tcp(host:port)/dbname?charset=utf8")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	var taskName string

	err=db.QueryRow("select taskName from tdb_indotask where id=?",846923).Scan(&taskName)
	if err != nil {
		panic(err)
	}
	fmt.Println(taskName)




}
