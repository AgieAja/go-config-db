package postgres

import (
	"database/sql"
	"fmt"

	//dont delete this package..driver for connection postgreSQL
	_ "github.com/lib/pq"
)

var SqlDb *sql.DB

//ConnPostgres - create connection database postgresSQL
func ConnPostgres(dbHost, dbUser, dbPass, dbName, dbTimeZone string, maxIdle, maxConn int) (*sql.DB, error) {
	// val := url.Values{}
	// val.Add("TimeZone", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	desc := getConnectionString(dbHost, dbUser, dbPass, dbName, dbTimeZone)
	sqlDb, err := sql.Open(`postgres`, desc)
	if err != nil {
		return nil, err
	}
	SqlDb = sqlDb

	errPing := sqlDb.Ping()
	if errPing != nil {
		return nil, errPing
	}

	sqlDb.SetMaxIdleConns(maxIdle)
	sqlDb.SetMaxOpenConns(maxConn)

	return sqlDb, nil
}

//getConnectionString - preparation connection database postgresSQL
func getConnectionString(dbHost, dbUser, dbPass, dbName, dbTimeZone string) string {
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbTimeZone)

	return desc
}
