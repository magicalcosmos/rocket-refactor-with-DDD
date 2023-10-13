package repositoryImpl

import (
	"fmt"
	"sync"

	"shareus.cn/domain/entity"
	r "shareus.cn/domain/repository"
)

type userRepositoryImpl struct {
	sync.Mutex
}

func NewUserRepositoryImplInstance() r.UserRepository {
	return &userRepositoryImpl{}
}

func (impl *userRepositoryImpl) Add(user entity.User) error {
	fmt.Println("username:", user.Username, "password:", user.Password)
	return nil
}

func (impl *userRepositoryImpl) Get() error {
	return nil
}

func (impl *userRepositoryImpl) Update() error {
	return nil
}

func (impl *userRepositoryImpl) Delete() error {
	return nil
}
