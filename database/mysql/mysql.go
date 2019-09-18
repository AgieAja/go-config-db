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

//InitConnMySQLDB - preparetion connection database mysql
func InitConnMySQLDB(dbHost,dbPort,dbUser,dbPass,dbName string) {
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	mysqldb, sqlErr = createConnMySQL(desc)
}

//GetMySQLDB - get connection db mysql
func GetMySQLDB() (*sql.DB, error) {
	return mysqldb, sqlErr
}

//createConnMySQL - create connection database mysql
func createConnMySQL(desc string) (*sql.DB, error) {
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

	mysqldb.SetMaxIdleConns(10)
	mysqldb.SetMaxOpenConns(10)

	return mysqldb, nil
}
