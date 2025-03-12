package main

import "github.com/wailsapp/wails/v3/pkg/application"

type GreetService struct{}

func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}

func (g *GreetService) OpenSettings() {
	app := application.Get()
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:            "Settings",
		URL:              "/settings",
		BackgroundColour: application.NewRGB(27, 38, 54),
	})
}
