package app

import "github.com/florinutz/go-donut-challenge/pkg/config"

// App is the logic behind the cli app.
type App struct {
	*config.Config
}

// New instantiates an App
func New(name string) *App {
	return &App{
		Config: config.New(name),
	}
}
