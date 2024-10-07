package entity

import (
	"github.com/google/uuid"
	"github.com/magicalcosmos/rocket/domain/vo"
)

type Function struct {
	ID       uuid.UUID   
	Location vo.Position 
}

func (f *Function) GetFunctionList() string {
	return ""
}
