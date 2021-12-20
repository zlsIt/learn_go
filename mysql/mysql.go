package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

// 使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql驱动） github.com/jmoiron/sqlx （基于mysql驱动的封装）
// go get github.com/go-sql-driver/mysql
// go get github.com/jmoiron/sqlx

var Db *sqlx.DB

type User struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CallPhone string    `db:"call_phone"`
	Status    int8      `db:"status"`
	CreateAt  time.Time `db:"createAt"`
	UpdateAt  time.Time `db:"updateAt"`
}

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:8066)/test?parseTime=true")
	if err != nil {
		panic(err)
	}
	Db = database
}

func main() {
	//install()
	//getAll()
	//update()
	//del()
	ts()
}

// 删除
func del() {
	sqlStr := "DELETE FROM `user` WHERE name = ?"
	res, err := Db.Exec(sqlStr, "tom")
	if err != nil {
		fmt.Println("del exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("delete succ: ", row)
}

// 更新
func update() {
	sqlStr := "UPDATE `test`.`user` SET `updateAt` = ? WHERE `id` = ?;"
	res, err := Db.Exec(sqlStr, time.Now(), "6e35cef0ea9742b18cc2f60d611260cb")
	if err != nil {
		fmt.Println("update failed err:", err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("update filed err:", err)
		return
	}
	fmt.Println("update rows:", affected)
}

// 根据条件获取多个值
func getAll() {
	var user []User
	sqlStr := "select `id`, `name`, `password`, `call_phone`, `status` , `createAt`, `updateAt` from user where name = ?"
	err := Db.Select(&user, sqlStr, "tom")
	if err != nil {
		fmt.Println("select data failed.", err)
		return
	}
	for _, u := range user {
		fmt.Println(u)
	}
}

// 新增
func install() {
	sqlStr := "INSERT INTO `user`(`id`, `name`, `password`, `call_phone`, `status`, `createAt`, `updateAt`) VALUES (?, ?, ?, ?, ?, ?, ?)"
	res, err := Db.Exec(sqlStr, strings.ReplaceAll(uuid.NewV1().String(), "-", ""), "tom", "123456", "110", 0, time.Now(), time.Now())
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows affected failed:", err)
		return
	}
	if affected == 0 {
		fmt.Println("install data failed.")
		return
	}
}

// 事务
func ts() {
	ts, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	sqlStr := "INSERT INTO `user`(`id`, `name`, `password`, `call_phone`, `status`, `createAt`, `updateAt`) VALUES (?, ?, ?, ?, ?, ?, ?)"
	res, err := ts.Exec(sqlStr, strings.ReplaceAll(uuid.NewV1().String(), "-", ""), "tom", "123456", "110", 0, time.Now(), time.Now())
	if err != nil {
		fmt.Println("exec failed, ", err)
		ts.Rollback()
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		ts.Rollback()
		return
	}
	fmt.Println("insert succ:", id)
	ts.Commit()
}
