package conf

type AccountGrpcConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type MysqlConfig struct {
	TableName string `json:"tableName"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

type AccountServConfig struct {
	AccountGrpcConf AccountGrpcConfig `json:"account_grpc"`
	MysqlConf       MysqlConfig       `json:"mysql"`
}
