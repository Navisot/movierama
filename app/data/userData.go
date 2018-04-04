package data

import (
	"errors"
	"fmt"

	"github.com/navisot/movierama/app"
)

func (s *Storage) User(id int) (*app.User, error) {

	u := new(app.User)

	fmt.Println(s)

	return u, errors.New("lol")

	err := s.DB.Where("user_id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}

	defer s.DB.Close()

	return u, nil
}

func (s *Storage) LoginUser(email, password string) (*app.User, error) {

	u := new(app.User)

	pwd := helpers.HashPassword(email, password)

	err := s.DB.Where("email = ? and password = ?", email, pwd).First(&u).Error
	if err != nil {
		return nil, err
	}

	defer s.DB.Close()

	return u, nil
}

func (s *Storage) RegisterUser(u *app.User) error {

	fmt.Println(u.Email)
	fmt.Println(u.Password)

	return errors.New("this is madness")
}
