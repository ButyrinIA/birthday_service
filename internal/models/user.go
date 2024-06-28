package models

import "time"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"-"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
	IsSubscribed bool   `json:"is_subscribed"`
}

func (u *User) GetBirthday() (*time.Time, error) {
	birthday, err := time.Parse("2006-01-02", u.Birthday)
	if err != nil {
		return nil, err
	}
	return &birthday, nil
}
