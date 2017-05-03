package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/t3reezhou/figure/figure"
	"github.com/t3reezhou/figure/figure/cfg"
)

var configFile = flag.String("config", "../../etc/figure.toml", "figure config file")

var _server *figure.Server

func reload() {}
func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(*configFile) == 0 {
		println("no config set")
		os.Exit(1)
	}
	cfg, err := cfg.ParseConfig(*configFile)
	if err != nil {
		println("Parse config:", err.Error())
		os.Exit(1)
	}

	_server, err = figure.NewServer(cfg)
	if err != nil {

	}

	_server.Run()

	_server.Close()
}
