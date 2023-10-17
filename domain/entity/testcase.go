package entity

import (
	"time"

	"github.com/google/uuid"
)

type TestCase struct {
	ID           uuid.UUID 
	TestCaseName string    
	CreateTime   time.Time 
	UpdateTime   time.Time 
	Status       bool     
}

func (c *TestCase) DeleteTestCase(id int64) bool {
	// TODO
	return true
}

func (c *TestCase) GetTestCase(id int64) string {
	return ""
}

func (c *TestCase) UpdateTestCase(id int64, testCase TestCase) bool {
	return true
}
