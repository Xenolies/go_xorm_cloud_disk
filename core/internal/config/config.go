package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		Driver     string
		DataSource string
	}
	Redis struct {
		DataSource string
		Password   string
		DB         int
	}
}
