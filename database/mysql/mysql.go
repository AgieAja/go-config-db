package mysql

import (
	"database/sql"
	"net/url"

	"fmt"

	//driver mysql database
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqldb *sql.DB
	sqlErr  error
)

//InitConnMySQLDB - preparation connection database mysql
func InitConnMySQLDB(dbHost, dbPort, dbUser, dbPass, dbName string, maxIdle, maxConn int) {
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	mysqldb, sqlErr = createConnMySQL(desc, maxIdle, maxConn)
}

//GetMySQLDB - get connection db mysql
func GetMySQLDB() (*sql.DB, error) {
	return mysqldb, sqlErr
}

//createConnMySQL - create connection database mysql
func createConnMySQL(desc string, maxIdle, maxConn int) (*sql.DB, error) {
	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	mysqldb, err := sql.Open(`mysql`, dsn)
	if err != nil {
		return nil, err
	}

	err = mysqldb.Ping()
	if err != nil {
		return nil, err
	}

	mysqldb.SetMaxIdleConns(maxIdle)
	mysqldb.SetMaxOpenConns(maxConn)

	return mysqldb, nil
}
