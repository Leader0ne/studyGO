package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go连接MySQL示例

func main() {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil { // dsn格式不正确时会报错
		fmt.Printf("dsn:%s invalid , err:%v\n", dsn, err)
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功！")

}
