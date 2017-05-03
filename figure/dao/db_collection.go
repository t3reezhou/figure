package dao

import (
	"time"

	"github.com/t3reezhou/figure/figure/cfg"
	"github.com/t3reezhou/figure/lib/mysql"

	"github.com/jmoiron/sqlx"
)

type DBCollection struct {
	Default *sqlx.DB
}

const (
	DefaultConnectWaiTimeout = 15 * time.Second
	DefaultCharset           = "utf8"
)

func NewDBCollection(config *cfg.Config) (*DBCollection, error) {
	dbcfg := config.DB
	db, err := mysql.OpenEx(dbcfg.Default.Host, dbcfg.Default.Port, dbcfg.User,
		dbcfg.Password, dbcfg.DBName, dbcfg.MaxIdleConns, dbcfg.MaxConns, DefaultConnectWaiTimeout, DefaultCharset)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	dbc := new(DBCollection)
	dbc.Default = db
	return dbc, nil
}
