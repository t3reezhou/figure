package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenEx(host string, port int, user string, password string, dbname string,
	maxIdleConns int, maxConnection int, connectWaitTimeout time.Duration, charset string) (*sqlx.DB, error) {
	db, err := openWrapDBEx(host, port, user, password, dbname, maxIdleConns, maxConnection, connectWaitTimeout, charset)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openWrapDBEx(host string, port int, user string, password string, dbname string,
	maxIdleConns int, maxConnection int, connectWaitTimeout time.Duration, charset string) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&charset=%s",
		user, password, host, port, dbname, connectWaitTimeout.String(), charset)
	sqlxDB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlxDB.SetMaxIdleConns(maxIdleConns)
	sqlxDB.SetMaxOpenConns(maxConnection)
	return sqlxDB, nil
}
