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
func createConnPostgresORM(desc string, maxIdle, maxConn int) (*gorm.DB, error) {
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

	sqlDbORM.DB().SetMaxIdleConns(maxIdle)
	sqlDbORM.DB().SetMaxOpenConns(maxConn)

	return sqlDbORM, nil
}

//createConnPostgres - create connection database postgresSQL
func createConnPostgres(desc string, maxIdle, maxConn int) (*sql.DB, error) {
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

	sqlDb.SetMaxIdleConns(maxIdle)
	sqlDb.SetMaxOpenConns(maxConn)

	return sqlDb, nil
}

//InitConnPostgresSQLDBORM - preparation connection database postgresSQL ORM
func InitConnPostgresSQLDBORM(dbHost, dbUser, dbPass, dbName,dbPort,dbSSL, dbTimeZone string, maxIdle, maxConn int) {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass,
		dbName,dbPort,dbSSL, dbTimeZone)

	sqlDbORM, sqlORMErr = createConnPostgresORM(desc, maxIdle, maxConn)
}

//InitConnPostgresSQLDB - preparation connection database postgresSQL ORM
func InitConnPostgresSQLDB(dbHost, dbUser, dbPass, dbName,dbPort,dbSSL, dbTimeZone string, maxIdle, maxConn int) {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass,
		dbName,dbPort,dbSSL, dbTimeZone)

	sqlDb, sqlErr = createConnPostgres(desc, maxIdle, maxConn)
}

//GetPostgresSQLDBORM - get connection db postgres ORM
func GetPostgresSQLDBORM() (*gorm.DB, error) {
	return sqlDbORM, sqlORMErr
}

//GetPostgresSQLDB - get connection db postgres
func GetPostgresSQLDB() (*sql.DB, error) {
	return sqlDb, sqlErr
}
