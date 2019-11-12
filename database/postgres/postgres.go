package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	//dont delete this package..driver for connection postgreSQL
	_ "github.com/lib/pq"
)

var (
	sqlDbORM          *gorm.DB
	sqlORMErr, sqlErr error
	sqlDb             *sql.DB
)

//createConnPostgresORM - create connection database postgresSQL
func createConnPostgresORM(desc string) (*gorm.DB, error) {
	// val := url.Values{}
	// val.Add("TimeZone", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	sqlDbORM, err := gorm.Open(`postgres`, desc)
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

//createConnPostgres - create connection database postgresSQL
func createConnPostgres(desc string) (*sql.DB, error) {
	// val := url.Values{}
	// val.Add("TimeZone", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	sqlDb, err := sql.Open(`postgres`, desc)
	if err != nil {
		return nil, err
	}

	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(10)

	return sqlDb, nil
}

//InitConnPostgresSQLDBORM - preparetion connection database postgresSQL ORM
func InitConnPostgresSQLDBORM(dbHost, dbUser, dbPass, dbName, dbTimeZone string) {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbTimeZone)

	sqlDbORM, sqlORMErr = createConnPostgresORM(desc)
}

//InitConnPostgresSQLDB - preparetion connection database postgresSQL ORM
func InitConnPostgresSQLDB(dbHost, dbUser, dbPass, dbName, dbTimeZone string) {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbTimeZone)

	sqlDb, sqlErr = createConnPostgres(desc)
}

//GetPostgresSQLDBORM - get connection db postgres ORM
func GetPostgresSQLDBORM() (*gorm.DB, error) {
	return sqlDbORM, sqlORMErr
}

//GetPostgresSQLDB - get connection db postgres
func GetPostgresSQLDB() (*sql.DB, error) {
	return sqlDb, sqlErr
}
