package postgres

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	//dont delete this package..driver for connection postgreSQL
	_ "github.com/lib/pq"
)

var (
	sqlDbORM  *gorm.DB
	sqlORMErr error
)

//createConnPostgresORM - create connection database postgresSQL
func createConnPostgresORM(desc string) (*gorm.DB, error) {
	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	sqlDbORM, err := gorm.Open(`postgres`, dsn)
	if err != nil {
		return nil, err
	}

	err = sqlDbORM.DB().Ping()
	if err != nil {
		return nil, err
	}

	sqlDbORM.DB().SetMaxIdleConns(10)
	sqlDbORM.DB().SetMaxOpenConns(10)

	return sqlDbORM, nil
}

//InitConnPostgresSQLDBORM - preparetion connection database postgresSQL
func InitConnPostgresSQLDBORM(dbHost, dbUser, dbPass, dbName string) {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbName)

	sqlDbORM, sqlORMErr = createConnPostgresORM(desc)
}

//GetPostgresSQLDBORM - get connection db postgres
func GetPostgresSQLDBORM() (*gorm.DB, error) {
	return sqlDbORM, sqlORMErr
}
