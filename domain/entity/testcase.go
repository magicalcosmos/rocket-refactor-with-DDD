package entity

import (
	"time"

	"github.com/google/uuid"
)

type TestCase struct {
	ID           uuid.UUID `json:"id"`
	TestCaseName string    `json:"test_case_name"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	Status       bool      `json:"status"`
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
