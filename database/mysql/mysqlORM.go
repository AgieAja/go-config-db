package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	mySQLORM *gorm.DB
	ormErr   error
)

//InitConnMySQLORM - preparation connection database mysql ORM
func InitConnMySQLORM(dbHost, dbPort, dbUser, dbPass, dbName string, maxIdle, maxConn int) {
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	mySQLORM, ormErr = createConnMySQLORM(desc, maxIdle, maxConn)
}

//GetMySQLORM - get connection ORM db mysql
func GetMySQLORM() (*gorm.DB, error) {
	return mySQLORM, ormErr
}

//createConnMySQLORM - create connection ORM database mysql
func createConnMySQLORM(desc string, maxIdle, maxConn int) (*gorm.DB, error) {
	mySQLORM, ormErr = gorm.Open(`mysql`, desc)
	if ormErr != nil {
		return nil, ormErr
	}

	ormErr = mySQLORM.DB().Ping()
	if ormErr != nil {
		return nil, ormErr
	}

	mySQLORM.DB().SetMaxIdleConns(maxIdle)
	mySQLORM.DB().SetMaxOpenConns(maxConn)
	return mySQLORM, nil
}
