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

func (e Environment) String() string {
	switch e {
	case Development:
		return "development"
	case Staging:
		return "staging"
	case Production:
		return "production"
	default:
		return "development"
	}
}

type EnvProvider interface {
	Get() Environment
}

type envProvider struct{}

func NewEnvProvider() EnvProvider {
	return &envProvider{}
}

func (ep envProvider) Get() Environment {
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
