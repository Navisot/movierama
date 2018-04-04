package data

import (
	"errors"
	"fmt"
	"github.com/navisot/movierama/models"
	"github.com/navisot/movierama/app/helpers"
)

func (s *Storage) User(id int) (*models.User, error) {

	u := new(models.User)

	fmt.Println(s)

	return u, errors.New("lol")

	err := s.DB.Where("user_id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}

	defer s.DB.Close()

	return u, nil
}

func (s *Storage) LoginUser(email, password string) (*models.User, error) {

	u := new(models.User)

	pwd := helpers.HashPassword(email, password)

	err := s.DB.Where("email = ? and password = ?", email, pwd).First(&u).Error
	if err != nil {
		return nil, err
	}

	defer s.DB.Close()

	return u, nil
}

func (s *Storage) RegisterUser(u *models.User) error {

	fmt.Println(u.Email)
	fmt.Println(u.Password)

	return errors.New("this is madness")
}
