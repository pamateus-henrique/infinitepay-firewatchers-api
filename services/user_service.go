// services/user_service.go
package services

import (
	"log"

	customErrors "github.com/pamateus-henrique/infinitepay-firewatchers-api/errors"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/repositories"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/utils"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/validators"
)

type UserService interface {
	Register(user *models.Register) error
	Login(login *models.Login) (*models.User, error)
	GetAllUsersPublicData() ([]*models.UserPublicData, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(user *models.Register) error {
	log.Println("Register: Starting user registration process")

	if err := validators.ValidateStruct(user); err != nil {
		log.Printf("Register: Validation error: %v", err)
		return &validators.ValidationError{Err: err}
	}

	log.Println("Register: Checking if user already exists")
	_, err := s.userRepo.GetUserByEmail(user.Email)
	if err == nil {
		log.Println("Register: User already exists")
		return &customErrors.AuthenticationError{Msg: "user already exists"}
	}

	log.Println("Register: Generating password hash")
	user.Password = utils.GeneratePassword(user.Password)

	log.Println("Register: Creating user")
	err = s.userRepo.CreateUser(user)
	if err != nil {
		log.Printf("Register: Error creating user: %v", err)
	} else {
		log.Println("Register: User created successfully")
	}
	return err
}

func (s *userService) Login(login *models.Login) (*models.User, error) {
	log.Println("Login: Starting login process")

	if err := validators.ValidateStruct(login); err != nil {
		log.Printf("Login: Validation error: %v", err)
		return nil, &validators.ValidationError{Err: err}
	}

	log.Println("Login: Retrieving user by email")
	user, err := s.userRepo.GetUserByEmail(login.Email)

	if err != nil {
		log.Println("Login: User not found or error retrieving user")
		return nil, &customErrors.AuthenticationError{Msg: "Invalid Email or password"}
	}

	log.Println("Login: Comparing passwords")
	if err := utils.ComparePassword(login.Password, user.Password); err != nil {
		log.Println("Login: Password comparison failed")
		return nil, &customErrors.AuthenticationError{Msg: "Invalid Email or password"}
	}

	log.Println("Login: Login successful")
	return user, nil
}

func (s *userService) GetAllUsersPublicData() ([]*models.UserPublicData, error) {
	log.Println("GetAllUsersPublicData: Starting retrieval of all users' public data")

	users, err := s.userRepo.GetAllUsersPublicData()
	if err != nil {
		log.Printf("GetAllUsersPublicData: Error retrieving users' public data: %v", err)
		return nil, err
	}

	log.Printf("GetAllUsersPublicData: Successfully retrieved public data for %d users", len(users))
	return users, nil
}
