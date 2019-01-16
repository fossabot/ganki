package user

type UserDTO struct {
	Username string
	Email    string
	Picture  string
}

type User struct {
	ID           string
	Username     string
	Email        string
	Picture      string
	PasswordHash string
}

type UserManager interface {
	Register(username, password string) error
	UpdateProfile(username string, password string, user User) error
	UpdatePassword(username string, password string, newPassword string) error
	LoginVerify(username, password string) error
	GetUserInfo(username string) (User, error)
}
