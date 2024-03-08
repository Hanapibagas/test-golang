package user

type UserCore struct {
	ID       uint   `json:"id"`
	Name     string `validate:"required"`
	Email    string `gorm:"default:null;unique"`
	Password string `validate:"min=6"`
	Role     string
}

type AuthCorePassword struct {
	ID       uint
	Password string
}

type UserDataInterface interface {
	Register(input UserCore) (data *UserCore, token string, err error)
	Login(email, password string) (data *UserCore, err error)
	UpdatePassword(id uint, input AuthCorePassword) error
	CheckPassword(savedPassword, inputPassword string) bool
	GetById(id uint) (*UserCore, error)
}

type UserServiceInterface interface {
	Register(input UserCore) (data *UserCore, token string, err error)
	Login(email, password string) (data *UserCore, token string, err error)
	UpdatePassword(id uint, input AuthCorePassword) error
	GetById(id uint) (*UserCore, error)
}
