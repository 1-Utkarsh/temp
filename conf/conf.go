package conf

import "sync"

type Config struct {
	DbAddress string
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
}

var conf *Config
var once sync.Once

func New() {
	once.Do(func() {
		conf = &Config{
			DbAddress: "localhost",
			DbPort:    "5432",
			DbUser:    "postgres",
			DbPass:    "password",
			DbName:    "postgres",
		}
	})
}

func Get() *Config {
	return conf
}