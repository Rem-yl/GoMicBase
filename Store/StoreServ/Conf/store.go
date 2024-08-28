package conf

type StoreSevConfig struct {
	MysqlConf MysqlConfig `json:"mysql"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DdName   string `json:"db_name"`
}
