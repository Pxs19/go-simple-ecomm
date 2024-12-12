package service

import (
	"errors"
	"log"

	"ecom2.5/internal/domain"
	"ecom2.5/internal/dto"
	"ecom2.5/internal/helper"
	"ecom2.5/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	// perform some db operation
	user, err := s.Repo.FindUser(email)

	// business logic

	return &user, err
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {

	hPass, err := s.Auth.CreateHashPassword(input.Password)

	if err != nil {
		return "", err
	}

	// log.Println(input)
	user, err := s.Repo.CreateUser(
		domain.User{
			Email:    input.Email,
			Password: hPass,
			Phone:    input.Phone,
		},
	)

	if err != nil {
		return "", err
	}

	// generate token
	log.Println(user)
	// userInfo := fmt.Sprintf("%v", user.ID, user.Email, user.UserType)
	token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
	println(token)

	return token, err

}

func (s UserService) Login(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	// compare password and generate token
	err = s.Auth.VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	// generate Token
	// token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType)

	// if err != nil {
	// 	return "", err
	// }

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)

}

func (s UserService) GetVerificationCode(input any) (int, error) {

	return 0, nil

}

func (s UserService) VerifyCode(id uint, code int) error {

	return nil

}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil

}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil

}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil

}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil

}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil

}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {

	return nil, nil

}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil

}

func (s UserService) GetOrder(u domain.User) ([]interface{}, error) {

	return nil, nil

}

func (s UserService) GetOrderById(id uint) ([]interface{}, error) {

	return nil, nil

}
