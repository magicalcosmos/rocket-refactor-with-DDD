package repository

import (
	"shareus.cn/domain/entity"
)

type UserRepository interface {
	Add(entity.User) error
	Update(entity.User) error
	Delete(id string) error
	Get() error
}
