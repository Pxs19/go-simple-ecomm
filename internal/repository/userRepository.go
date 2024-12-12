package repository

import (
	"errors"
	"log"

	"ecom2.5/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)

	// more function will come. . .
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur userRepository) CreateUser(usr domain.User) (domain.User, error) {

	err := ur.db.Create(&usr).Error

	if err != nil {
		log.Printf("create user error", err)
		return domain.User{}, errors.New("failed to create user")
	}

	println("this isL", err)

	return usr, nil
}

func (ur userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	// err := ur.db.First(&user, "email = ?", email)
	// or
	err := ur.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		log.Printf("find user by email error", err)
		return domain.User{}, errors.New("failed to find user")
	}

	return user, nil
}

func (ur userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User

	// err := ur.db.First(&user, "email = ?", email)
	// or
	err := ur.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		log.Printf("find user by email error", err)
		return domain.User{}, errors.New("failed to find user")
	}

	return user, nil
}

func (ur userRepository) UpdateUser(id uint, usr domain.User) (domain.User, error) {

	var user domain.User

	err := ur.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(usr).Error

	if err != nil {
		log.Printf("update user error", err)
		return domain.User{}, errors.New("update user user")
	}

	return user, nil
}
