package executor

import (
	"github.com/google/uuid"
	"shareus.cn/domain/entity"
	"shareus.cn/domain/repository"
	"shareus.cn/infrastructure-mongo/repositoryImpl"
)

var userRepositoryDomain repository.UserRepository

func init() {
	impl := repositoryImpl.NewUserRepositoryImplInstance()
	userRepositoryDomain = impl
}

func Save(username string, password string) {
	userRepositoryDomain.Add(entity.User{
		ID:       uuid.New(),
		Username: username,
		Password: password,
	})
}

func Update(id uuid.UUID, username string, password string) {
	userRepositoryDomain.Update(entity.User{
		ID:       id,
		Username: username,
		Password: password,
	})
}

func Delete(id string) {
	userRepositoryDomain.Delete(id)
}

func Get(username string, password string) {
	userRepositoryDomain.Get()
}
