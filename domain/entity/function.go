package entity

import (
	"github.com/google/uuid"
	"shareus.cn/domain/vo"
)

type Function struct {
	ID       uuid.UUID   `json:"id"`
	Location vo.Position `json:"location"`
}

func (f *Function) GetFunction() string {
	return ""
}
