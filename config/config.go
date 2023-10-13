// Package config handles environment variables.
package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config contains environment variables.
type Config struct {
	Port               string  `envconfig:"PORT" default:"8000"`
	JaegerAgentHost    string  `envconfig:"JAEGER_AGENT_HOST" default:"localhost"`
	JaegerAgentPort    string  `envconfig:"JAEGER_AGENT_PORT" default:"6831"`
	JaegerSamplerType  string  `envconfig:"JAEGER_SAMPLER_TYPE" default:"const"`
	JaegerSamplerParam float64 `envconfig:"JAEGER_SAMPLER_PARAM" default:"1"`
	ExampleAPIEndpoint string  `envconfig:"EXAMPLE_API_ENDPOINT" required:"true"`
}

// LoadConfig reads environment variables and populates Config.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	var c Config

	err := envconfig.Process("", &c)

	return &c, err
}
