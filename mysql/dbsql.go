package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err := sql.Open("mysql", "root:ba199035@tcp(127.0.0.1:3306)/e1?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()

	//查询数据，指定字段名，返回sql.Rows结果集
	rows, _ := db.Query("select id,auth_name,status from account_auth")
	id := 0
	auth_name := ""
	status := ""
	for rows.Next() {
		rows.Scan(&id, &auth_name, &status)
		fmt.Println(id, auth_name, status)
	}

}
