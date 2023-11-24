package env

import (
	"os"
)

type Environment int

const (
	Development Environment = iota + 1
	Staging
	Production
)

type EnvProvider interface {
	Get() Environment
}

type envProvider struct{}

func NewEnvProvider() EnvProvider {
	return &envProvider{}
}

func (e envProvider) Get() Environment {
	switch os.Getenv("ENV") {
	case "development":
		return Development
	case "staging":
		return Staging
	case "production":
		return Production
	default:
		return Development
	}
}
