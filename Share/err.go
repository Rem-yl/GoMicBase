package share

// err
const (
	ErrDatabaseConn       = "Database Connect Error"
	ErrDatabaseInit       = "Database Init Error"
	ErrConfigNotFound     = "Config File Not Found"
	ErrConfigReadFailed   = "Read Config File Failed"
	ErrListen             = "Listen Failed"
	ErrGrpcServerFailed   = "Grpc Server Failed"
	ErrCreateAccount      = "Create Account Failed"
	ErrAccountNotFound    = "Account Not Found"
	ErrConfigFileNotFound = "Config File Not Found"
	ErrGrpcDialFailed     = "Grpc Dial Failed"
	ErrInvalidToken       = "Invalid Token"
	ErrParseAccount       = "Parse Account Failed"
	ErrNotRegister        = "Register First"
	ErrGenJWTFailed       = "Generate JWT Failed"
	ErrGetConsulClient    = "Get Consul Client Failed"
	ErrGrpcRegister       = "Grpc Registe Failed"
	ErrWebRegister        = "Web Registe Failed"
	AccountExisted        = "Account Existed"
	BrandExisted          = "Brand Existed"
	ErrBrandNotFound      = "Brand Not Found"
)
