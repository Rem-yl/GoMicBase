package conf

type AccountConfig struct {
	AccountServConf AccountServConfig `json:"account_serv"`
	AccountWebConf  AccountWebConfig  `json:"account_web"`
	MysqlConf       MysqlConfig       `json:"mysql"`
	ConsulConf      ConsulConfig      `json:"consul"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type AccountServConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

type AccountWebConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

type MysqlConfig struct {
	TableName string `json:"tableName"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
}
