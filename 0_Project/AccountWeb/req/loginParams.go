package req

type LoginByPassword struct {
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone" binding:"required"`
}
