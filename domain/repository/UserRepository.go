package repository

import (
	"shareus.cn/domain/entity"
)

type UserRepository interface {
	Add(entity.User) error
	Update() error
	Delete() error
	Get() error
}
