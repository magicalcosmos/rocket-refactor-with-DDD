package entity

type UserEntity struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Password int    `json:"password"`
}
