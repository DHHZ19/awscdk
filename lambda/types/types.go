package types

import (
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username  string `json:"username"`
	Passsword string `json:"password"`
}

type User struct {
	Username     string `json:"username"`
	PasswrodHash string `json:"password"`
}

func NewUser(registerUser RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Passsword), 10)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:     registerUser.Username,
		PasswrodHash: string(hashedPassword),
	}, nil
}

func ValidatePassword(hashedPassword, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil
}
