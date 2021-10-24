package main

// 实现错误处理机制

import (
	"fmt"
	"gee"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println("open mysql success")
	Db = database
	// defer Db.Close()
}

func main() {
	r := gee.New() // engine
	// initDb()
	// r.GET("/", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, gee.H{
	// 		"name": "zsd",
	// 	})
	// })
	// insertPerson("yzf", "1", "fans@ww.com")
	// selectPerson(2)
	// updatePerson("wkm", 2)
	// deletePerson(4)
	// 事务
	testss()
	r.Run(":9999")
}

// init Table
func initDb() {
	var schema1 = `
CREATE TABLE person (
	user_id int(11) NOT NULL AUTO_INCREMENT,
	username varchar(260) DEFAULT NULL,
	sex varchar(260) DEFAULT NULL,
	email varchar(260) DEFAULT NULL,
	PRIMARY KEY (user_id)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
`

	var shema2 = `
CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
`
	Db.MustExec(schema1)
	Db.MustExec(shema2)
	fmt.Println("mysql success")
}

// insert
func insertPerson(username, sex, email string) {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", username, sex, email)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert succ:", id)
}

// select
func selectPerson(id int) {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", id)

	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select succ:", person)
}

// Update
func updatePerson(newName, userid interface{}) {
	res, err := Db.Exec("update person set username=? where user_id=?", newName, userid)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

// Delete
func deletePerson(userid interface{}) {
	res, err := Db.Exec("delete from person where user_id=?", userid)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

// MySQL事务
// Db.Begin()        开始事务
// Db.Commit()        提交事务
// Db.Rollback()     回滚事务
func testss() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	// insert 1
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "YY", "2", "yy@666.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	// insert 2
	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "YY", 1, 123)
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}
