package user

import (
	"fmt"

	"github.com/dulev/ganki/server/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

type UserService interface {
	RegisterUser(user models.User) error
	Authenticate(username, password string) error
}

func NewUserService(database *gorm.DB) UserService {
	return &UserServiceImpl{
		Database: database,
	}
}

type UserServiceImpl struct {
	Database *gorm.DB
}

// TODO: Maybe later add confirmation email.
func (us *UserServiceImpl) RegisterUser(user models.User) error {
	fmt.Printf("Trying to register user: %#v\n", user)

	passwordHash, err := HashPassword(user.Password)
	if err != nil {
		// TODO:
		return errors.Wrap(err, "couldn't hash password")
	}
	user.Password = passwordHash

	if err := us.Database.Create(&user).Error; err != nil {
		fmt.Printf("Error writing to database: %v\n", err)
	}

	return nil
}

func (us *UserServiceImpl) Authenticate(username, password string) error {
	var user models.User
	if err := us.Database.Where("username = ?", username).First(&user).Error; err != nil {
		return errors.Wrap(err, "couldn't get user from database")
	}

	hash := user.Password
	if !HashPasswordMatch(hash, password) {
		return errors.New("password doesn't match against hash")
	}

	return nil
}

// func UpdateProfile(username string, user User) error
// func UpdatePassword(username string, password string, newPassword string) error
// func LoginVerify(username, password string) error
// func GetUserInfo(username string) (User, error)

func HashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashBytes), err
}

func HashPasswordMatch(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
