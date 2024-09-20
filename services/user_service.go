// services/user_service.go
package services

import (
	customErrors "github.com/pamateus-henrique/infinitepay-firewatchers-api/errors"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/utils"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/validators"
)

type UserService interface {
    Register(user *models.Register) error
    Login(login *models.Login) (*models.User, error)
    // GetUser(id int) (*models.User, error)
}

type userService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}

func (s *userService) Register(user *models.Register) error {

    if err := validators.ValidateStruct(user); err != nil {
        return &validators.ValidationError{Err: err}
    }

    // Check if user already exists
    _, err := s.userRepo.GetUserByEmail(user.Email)
    if err == nil {
        return &customErrors.AuthenticationError{Msg: "user already exists"}
    }

    user.Password = utils.GeneratePassword(user.Password)
	
    // Create user
    return s.userRepo.CreateUser(user)
}

func (s *userService) Login(login *models.Login) (*models.User, error) {

    if err := validators.ValidateStruct(login); err != nil {
        return nil, &validators.ValidationError{Err: err}
    }

    user, err := s.userRepo.GetUserByEmail(login.Email)

    if err != nil {
        return nil, &customErrors.AuthenticationError{Msg: "Invalid Email or password"}
    }

    if err := utils.ComparePassword(login.Password, user.Password); err != nil {
		return nil, &customErrors.AuthenticationError{Msg: "Invalid Email or password"}
	}


    return user, nil
}

// func (s *userService) GetUser(id int) (*models.User, error) {
//     return s.userRepo.GetUserByID(id)

// }
