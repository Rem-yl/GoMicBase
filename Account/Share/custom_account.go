package share

type CustomAccount struct {
	Id             uint32 `json:"id"`
	Name           string `json:"string"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	Salt           string `json:"salt"`
	HashedPassword string `json:"hashed_password"`
}
