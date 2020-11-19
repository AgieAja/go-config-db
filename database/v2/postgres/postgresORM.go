package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var SqlDbORM *gorm.DB

//ConnPostgresORM - create connection database postgresSQL ORM
func ConnPostgresORM(dbHost, dbUser, dbPass, dbName,dbPort,dbSSL, dbTimeZone string, maxIdle, maxConn int) (*gorm.DB, error) {
	// val := url.Values{}
	// val.Add("TimeZone", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s&%s", desc, val.Encode())

	desc := getConnectionStringORM(dbHost, dbUser, dbPass, dbName,dbPort,dbSSL, dbTimeZone)
	sqlDbORM, err := gorm.Open(`postgres`, desc)
	if err != nil {
		return nil, err
	}
	SqlDbORM = sqlDbORM

	errPing := sqlDbORM.DB().Ping()
	if errPing != nil {
		return nil, errPing
	}

	sqlDbORM.DB().SetMaxIdleConns(maxIdle)
	sqlDbORM.DB().SetMaxOpenConns(maxConn)

	return sqlDbORM, nil
}

//getConnectionStringORM - preparation connection database postgresSQL ORM
func getConnectionStringORM(dbHost, dbUser, dbPass, dbName,dbPort,dbSSL, dbTimeZone string) string {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser,
		dbPass, dbName,dbPort,dbSSL, dbTimeZone)

	return desc
}
