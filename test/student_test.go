package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {

	db := DBConnect{
		DBUrl:"root:root1243@tcp(127.0.0.1:3306)/test?charset=utf8mb4",
	}
	var err error
	db.DB, err = sql.Open("mysql", db.DBUrl)
	if err != nil {
		fmt.Printf("sql connect error : %s", err)
		return
	}
	fmt.Println(db.DBUrl)
	fmt.Println(db.DB.Stats())
	db.queryStudent()

}
