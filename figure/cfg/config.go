package cfg

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HttpAddr   string    `toml:"http_addr"`
	UnixSocket string    `toml:"unix_socket"`
	DB         *dbConfig `toml:"database"`
}

type dbConfig struct {
	Default *Addr `toml:"default"`

	User         string `toml:"user"`
	Password     string `toml:"password"`
	DBName       string `toml:"db_name"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxConns     int    `toml:"max_conns"`
	SqlLogPath   string `toml:"sql_log_path"`
	OpenSqlLog   bool   `toml:"open_sql_log"`

	SlowLogTimeout int64 `toml:"slow_log_timeout"`

	Partitions int  `toml:"partitions"` // 分表量
	ReadOnly   bool `toml:"read_only"`
}

type Addr struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func ParseConfig(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}
	cfg := new(Config)

	_, err = toml.Decode(string(data), cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
