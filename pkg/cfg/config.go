package cfg

type MysqlConfig struct {
	TableName string `json:"tableName"`
	Host      string `json:"host"`
	Port      int32  `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

type GrpcServConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}
