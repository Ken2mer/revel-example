package models

import "fmt"

type User struct {
	UserId int
	Name   string
}

func (u *User) String() string {
	return fmt.Sprintf("User id: %d, name: %s", u.UserId, u.Name)
}
