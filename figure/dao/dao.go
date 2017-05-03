package dao

import "github.com/t3reezhou/figure/figure/cfg"

var (
	FigureDaoManager *FigureDao
)

type Dao struct{}

func NewDao(config *cfg.Config) error {
	dbc, err := NewDBCollection(config)
	if err != nil {
		panic(err)
	}
	FigureDaoManager = NewFigureDaoManager(dbc.Default)
	return nil
}
