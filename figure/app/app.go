package app

import (
	"github.com/t3reezhou/figure/figure/bll"
	"github.com/t3reezhou/figure/figure/cfg"
	"github.com/t3reezhou/figure/figure/dao"
)

type App struct {
	*dao.Dao
	*bll.Bll
	*cfg.Config
}

func NewApp(config *cfg.Config) (*App, error) {
	a := new(App)
	a.Config = config
	if err := dao.NewDao(config); err != nil {
		return nil, err
	}
	return a, nil
}
