package executor

import (
	"github.com/google/uuid"
	"shareus.cn/domain/entity"
	"shareus.cn/domain/repository"
	"shareus.cn/infrastructure-mongo/repositoryImpl"
)

func Save(username string, password string) {
	impl := repositoryImpl.NewUserRepositoryImplInstance()
	var userRepositoryDomain repository.UserRepository = impl
	userRepositoryDomain.Add(entity.User{
		ID:       uuid.UUID{},
		Username: username,
		Password: password,
	})
}
