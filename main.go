package main

import (
	"embed"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {}

func main() {

	app := application.New(application.Options{
		Name:        "JSON Editor",
		Description: "A demo of using raw HTML & CSS",
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:              "JSON Editor",
		Width:              1000,
		Height:             618,
		DisableResize:      false,
		ZoomControlEnabled: false,
		BackgroundColour:   application.NewRGB(6, 7, 15),
		URL:                "/",
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
