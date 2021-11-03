package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const databaseNameSource = "root:mjsbqsn741@tcp(127.0.0.1:3306)/ea000001"

type parent struct {
	id    int
	name  string
	phone string
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB faild, err: %v\n", err)
		return
	}
	fmt.Printf("连接数据库成功\n")
	queryOne(1)
	queryRaws(0)
	//insert("咲", "13532565342")
	//update("森咲", "18888889999", 4)
	//delete(8)
}

func initDB() (err error) {
	//	数据库信息
	// 连接数据库
	db, err = sql.Open("mysql", databaseNameSource) // 不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	fmt.Println("连接数据库成功!")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

// 查询
func queryOne(n int) {
	var p1 parent
	// 1. 写查询单条记录的sql语句
	sqlStr := `select id, name, phone from parent where id=?`
	// 2. 执行
	instance := db.QueryRow(sqlStr, n)
	// 3. 拿到结果
	_ = instance.Scan(&p1.id, &p1.name, &p1.phone)
	fmt.Printf("p1: %#v\n", p1)
	fmt.Println("=========================")
}

// 查询多行
func queryRaws(n int) {
	var p1 parent
	sqlStr := `select * from parent where id>?`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query faild, err: %v\n", sqlStr, err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&p1.id, &p1.name, &p1.phone)
		fmt.Printf("%#v\n", p1)
	}
	fmt.Println("=========================")
}

// 插入
func insert(name, phone string) {
	sqlStr := `insert into parent(name, phone) values(?, ?)`
	ret, err := db.Exec(sqlStr, name, phone)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsertId failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theId)
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed, err: %d\n", err)
		return
	}
	fmt.Printf("rowsAffected is: %d\n", n)
	fmt.Println("=========================")
}

// 更新
func update(name, phone string, id int) {
	sqlStr := `update parent set name=?, phone=? where id=?`
	ret, err := db.Exec(sqlStr, name, phone, id)
	if err != nil {
		fmt.Printf("uddate failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf(" get rowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
	fmt.Println("=========================")
}

// 删除
func delete(id int) {
	sqlStr := `delete from parent where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete parent failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed, err: %v\n", n)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
	fmt.Println("=========================")
}
