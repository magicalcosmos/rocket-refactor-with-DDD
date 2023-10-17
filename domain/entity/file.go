package entity

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID         uuid.UUID 
	FileName   string   
	FilePath   string  
	Status     bool   
	CreateTime time.Time 
	UpdateTime time.Time
	CreateUser User    
}

func (f *File) uploadFiles(file *File) bool {
	return true
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
