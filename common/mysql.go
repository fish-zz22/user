package common

import (

	"github.com/micro/go-micro/v2/config"
)

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	DataBase string `json:"data_base"`
	Port int64 `json:"port"`
}

func GetMysqlConfig(config config.Config,path ...string) *MysqlConfig {
	mySqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mySqlConfig)

	return mySqlConfig
}
