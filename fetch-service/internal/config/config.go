package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// DefaultConfig :nodoc:
type DefaultConfig struct {
	Port      string `json:"8081"`
	JwtSecret string `json:"jwtSecret"`
}

var (
	instance      *DefaultConfig
	once          sync.Once
	viperInstance *viper.Viper
	viperInit     sync.Once
)

// New :nodoc:
func New() *DefaultConfig {
	v := Viper()
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("config file not found: ", err)
	}

	err := v.Unmarshal(&instance)
	if err != nil {
		panic(err)
	}

	return instance
}

func Viper() *viper.Viper {
	viperInit.Do(func() {
		viperInstance = viper.NewWithOptions(viper.KeyDelimiter("::"))
		viperInstance.AddConfigPath("./config")
		viperInstance.SetConfigName("config")
	})
	return viperInstance
}

func getConfig() *DefaultConfig {
	if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
		return New()
	}

	once.Do(func() {
		if instance == nil {
			instance = New()
		}
	})
	return instance
}

// AppAddress :nodoc:
func (c *DefaultConfig) AppAddress() string {
	return fmt.Sprintf(":%v", c.Port)
}

// AppAddress :nodoc:
func GetConfig() *DefaultConfig {
	return getConfig()
}
