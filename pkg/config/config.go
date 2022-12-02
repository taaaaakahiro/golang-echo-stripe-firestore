package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type envConfig struct {
	Port       int    `env:"PORT"`
	ProjectID  string `env:"PROJECT_ID"`
	Collection string `env:"COLLECTION"`
	Document   string `env:"DOCUMENT"`
	StripeKey  string `env:"STRIPE_SECRET_KEY"`
}

func LoadConfig(ctx context.Context) (*envConfig, error) {
	var goenv envConfig
	err := envconfig.Process(ctx, &goenv)
	if err != nil {
		log.Fatal(err)
	}
	return &goenv, nil

}
