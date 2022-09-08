package models

import "golang.org/x/crypto/bcrypt"

const (
	UsersCollection = "users"
)

type User struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// return string(hashedPassword) == u.Password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
