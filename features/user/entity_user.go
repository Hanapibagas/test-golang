package user

type UserCore struct {
	ID       uint `json:"id"`
	Name     string
	Email    string `gorm:"default:null;unique"`
	Password string
	Role     string
}

type UserDataInterface interface {
	Register(input UserCore) (data *UserCore, token string, err error)
	Login(email, password string) (data *UserCore, err error)
}

type UserServiceInterface interface {
	Register(input UserCore) (data *UserCore, token string, err error)
	Login(email, password string) (data *UserCore, token string, err error)
}
