package infrastructure

import (
	"fmt"
	"sync"

	"github.com/magicalcosmos/rocket/domain/entity"
	r "github.com/magicalcosmos/rocket/domain/repository"
)

type userRepositoryImpl struct {
	sync.Mutex
}

func NewUserRepositoryImplInstance() r.IUserRepository {
	return &userRepositoryImpl{}
}

func (impl *userRepositoryImpl) Add(user entity.User) error {
	fmt.Println("username:", user.Username, "password:", user.Password, "id: ", user.ID)
	return nil
}

func (impl *userRepositoryImpl) Update(user entity.User) error {
	return nil
}

func (impl *userRepositoryImpl) Delete(id string) error {
	return nil
}

func (impl *userRepositoryImpl) Get() error {
	return nil
}
