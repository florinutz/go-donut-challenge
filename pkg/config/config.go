package config

import "github.com/florinutz/go-donut-challenge/pkg/container"

// Config holds flags, env vars and config files contents
type Config struct {
	*container.Container
	name string
}

func (c Config) Name() string {
	return c.name
}

func New(name string) *Config {
	return &Config{
		Container: container.New(),
		name:      name,
	}
}
