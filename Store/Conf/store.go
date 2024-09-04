package conf

import share "github.com/GoMicBase/Share"

type StoreConfig struct {
	StoreServConf StoreServConfig    `json:"store_serv"`
	StoreWebConf  StoreWebConfig     `json:"store_web"`
	MysqlConf     MysqlConfig        `json:"mysql"`
	ConsulConf    share.ConsulConfig `json:"consul"`
}

type StoreServConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int32  `json:"port"`
	Id   string `json:"id"`
}

type StoreWebConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int32  `json:"port"`
	Id   string `json:"id"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DdName   string `json:"db_name"`
}
