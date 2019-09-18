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
func InitConnMySQLORM(dbHost,dbPort,dbUser,dbPass,dbName string) {
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	mySQLORM, ormErr = createConnMySQLORM(desc)
}

//GetMySQLORM - get connection ORM db mysql
func GetMySQLORM() (*gorm.DB, error) {
	return mySQLORM, ormErr
}

//createConnMySQLORM - create connection ORM database mysql
func createConnMySQLORM(desc string) (*gorm.DB, error) {
	mySQLORM, ormErr = gorm.Open(`mysql`, desc)
	if ormErr != nil {
		return nil, ormErr
	}

	ormErr = mySQLORM.DB().Ping()
	if ormErr != nil {
		return nil, ormErr
	}

	mySQLORM.DB().SetMaxIdleConns(10)
	mySQLORM.DB().SetMaxOpenConns(10)
	return mySQLORM, nil
}
