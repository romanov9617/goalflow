package config

import (
	"encoding/json"
	"fmt"
	"io"
)

type Config struct {
	API      APIConfig     `json:"api"`
	Postgres PostgreConfig `json:"postgres"`
}

type PostgreConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	SSLMode  string `json:"ssl_mode"`
}

type APIConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig(r io.Reader) (*Config, error) {
	var cfg Config

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("LoadConfig: %w", err)
		type T string
	}
	return &cfg, nil
}
