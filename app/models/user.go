package models

import "fmt"

type User struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

func (u *User) String() string {
	return fmt.Sprintf("User id: %d, name: %s", u.UserId, u.Name)
}
