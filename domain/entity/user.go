package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func (u *User) MatchPwd(pwd string) bool {
	return true
}
