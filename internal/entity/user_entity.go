package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type GetUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return errors.New("Required fields must be filled")
	}
	return nil
}

// HashPassword hashes the user's password and replaces the plain text password
func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}
