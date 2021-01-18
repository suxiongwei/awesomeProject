package main

import (
	"database/sql"
	"fmt"
	_"github.com/Go-SQL-Driver/MySQL"
)


// 测试go数据库编程

/**
数据库表结构
create table user_info(
	`urid` int not null auto_increment,
	`username` varchar(64) default null,
	`departname` varchar(64) default null,
	`created` date default null,
	primary key(urid)
);
 */

// 定义数据库连接信息
type DbConn struct {
	Dsn string// 数据库驱动连接字符串
	Db *sql.DB
}

// user_info的映射对象
type UserTable struct {
	Uid int
	UserName string
	Department string
	Created string
}

func execData(dbConn *DbConn)  {
	count,id,err := dbConn.ExecData("INSERT user_info(username,departname,created) VALUES ('John','business Group','2018-07-03')")
	//count,id,err := dbConn.ExecData("UPDATE user_info SET created = '2019-07-03' WHERE username = 'John'")
	if err != nil{
		fmt.Println(err.Error())
	}else {
		fmt.Println("受影响行数", count)
		fmt.Println("自增id", id)
	}
}

// 封装增删改数据的函数，该函数直接使用DB的Exec()方法实现数据操作
func (dbconn *DbConn) ExecData(sqlString string) (count, id int64, err error) {
	result,err := dbconn.Db.Exec(sqlString)
	if err != nil {
		panic(err)
		return
	}
	if id,err = result.LastInsertId(); err != nil {
		panic(err)
		return
	}
	if count,err = result.RowsAffected(); err != nil{
		panic(err)
		return
	}
	return count, id, err
}

// 测试封装的PreExecData()
func preExecData(dbConn *DbConn) {
	count, id, err := dbConn.PreExecData("INSERT user_info(username,departname,created) VALUES (?,?,?)",
		"suxiongwei", "CUMT", "2014-09-02")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("受影响行数", count)
		fmt.Println("自增id", id)
	}
}

// 封装增删改查的函数，该函数使用预编译语句加Exec实现增删改
func (dbConn *DbConn) PreExecData(sqlString string, args ...interface{})(count, id int64, err error)  {
	// 获得预编译对象stmt
	stmt, err := dbConn.Db.Prepare(sqlString)
	defer stmt.Close()
	if err != nil {
		panic(err)
		return
	}

	result, err := stmt.Exec(args...)

	if err != nil {
		panic(err)
		return
	}

	if id,err = result.LastInsertId(); err != nil{
		panic(err)
		return
	}

	if count,err = result.RowsAffected(); err != nil {
		panic(err)
		return
	}
	return count, id, nil
}

func main3()  {
	var err error
	dbConn := DbConn{
		Dsn: "root:12345678@tcp(127.0.0.1:3306)/xingluo_ods?charset=utf8",
	}
	dbConn.Db,err = sql.Open("mysql", dbConn.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbConn.Db.Close()

	// 1.测试封装的execData
	//execData(&dbConn)

	// 2.测试封装的PreExecData()
	//preExecData(&dbConn)

	// 3.测试单行数据


}


