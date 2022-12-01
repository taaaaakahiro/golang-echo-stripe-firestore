package config

import "os"

type FSConfig struct {
	ProjectID  string
	Collection string
	Document   string
}

func NewFSConfig() *FSConfig {
	return &FSConfig{
		ProjectID:  os.Getenv("PROJECT_ID"),
		Collection: os.Getenv("COLLECTION"),
		Document:   os.Getenv("DOCUMENT"),
	}

}
