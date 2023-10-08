package entity

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID         uuid.UUID `json:"id"`
	FileName   string    `json:"file_name"`
	FilePath   string    `json:"file_path"`
	Status     bool      `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateUser User      `json:"create_user"`
}

func (f *File) ReplaceFiles(file *File) bool {
	return true
}

func (f *File) UpdateFile(sourceCode string) bool {
	return true
}

func (f *File) Download() string {
	return ""
}
