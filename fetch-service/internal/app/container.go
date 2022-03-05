package app

import (
	"github.com/bagusays/efishery-technical-test/internal/config"
	"github.com/bagusays/efishery-technical-test/internal/service"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/currency_converter"
	"github.com/bagusays/efishery-technical-test/internal/shared/integration/efishery"
)

// Container :nodoc:
type Container struct {
	Cfg             *config.DefaultConfig
	AuthService     *service.Auth
	ResourceService *service.Resource
}

// NewContainer :nodoc:
func NewContainer(cfg *config.DefaultConfig) *Container {
	currencyConverterClient := currency_converter.NewClient()
	efisheryClient := efishery.NewClient()
	return &Container{
		AuthService:     service.NewAuth(),
		ResourceService: service.NewResource(currencyConverterClient, efisheryClient),
		Cfg:             cfg,
	}
}
