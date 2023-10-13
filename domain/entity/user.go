package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}

func (u *User) MatchPwd(pwd string) bool {
	return true
}
