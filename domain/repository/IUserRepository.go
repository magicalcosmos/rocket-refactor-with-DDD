package domain

import (
	"github.com/magicalcosmos/rocket/domain/entity"
)

type IUserRepository interface {
	Add(entity.User) error
	Update(entity.User) error
	Delete(id string) error
	Get() error
}
