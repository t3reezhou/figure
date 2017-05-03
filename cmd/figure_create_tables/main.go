package main

import (
	"flag"
	"os"

	"github.com/t3reezhou/figure/figure/cfg"

	"github.com/t3reezhou/figure/figure/dao"
)

func createTables() error {
	var err error

	println("Create figure")
	err = dao.FigureDaoManager.CreateTable()
	if err != nil {
		return err
	}
	return nil
}

var configFile = flag.String("config", "../etc/figure.toml", "figure config file")

func main() {
	flag.Parse()
	//
	cfg, err := cfg.ParseConfig(*configFile)
	if err != nil {
		panic(err)
	}

	if err := dao.NewDao(cfg); err != nil {
		println("Dao new failed:", err.Error())
		os.Exit(1)
	}

	if err := createTables(); err != nil {
		println("FAILED: ", err.Error())
		os.Exit(1)
	}
}
