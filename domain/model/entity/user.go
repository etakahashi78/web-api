package entity

import (
	"errors"
	"time"
)

type User struct {
	ID        int64
	Email     string
	Password  string // ハッシュ化されたパスワードを格納することを想定
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email string, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	now := time.Now()
	// ここでパスワードのハッシュ化などの処理を行うことが考えられます
	return &User{
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (u *User) ChangePassword(newPassword string) error {
	if newPassword == "" {
		return errors.New("password cannot be empty")
	}
	// ここで新しいパスワードのハッシュ化などの処理を行うことが考えられます
	u.Password = newPassword
	u.UpdatedAt = time.Now()
	return nil
}
