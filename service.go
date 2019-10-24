package service

import (
	"github.com/glvd/conversion"
	"github.com/goextension/tool"
)

// Service ...
type Service struct {
	secret string
	task   *conversion.Task
}

// NewService ...
func NewService() *Service {
	secret := tool.GenerateRandomString(32)
	task := initConversion()
	return &Service{
		secret: secret,
		task:   task,
	}
}
