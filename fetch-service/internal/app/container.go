package app

import (
	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/service"
)

// Container :nodoc:
type Container struct {
	Cfg         *config.DefaultConfig
	AuthService *service.Auth
}

// NewContainer :nodoc:
func NewContainer(cfg *config.DefaultConfig) *Container {
	return &Container{
		AuthService: service.NewAuth(),
		Cfg:         cfg,
	}
}
