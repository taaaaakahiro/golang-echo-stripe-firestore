package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type EnvConfig struct {
	Port       int    `env:"PORT"`
	StripeKey  string `env:"STRIPE_SECRET_KEY"`
	ProjectID  string `env:"FS_PJ_ID"`
	Collection string `env:"FS_COLLECTION"`
	Document   string `env:"FS_SUBSC_DOCUMENT"`
}

func LoadConfig(ctx context.Context) (*EnvConfig, error) {
	var goenv EnvConfig
	err := envconfig.Process(ctx, &goenv)
	if err != nil {
		log.Fatal(err)
	}
	return &goenv, nil

}
