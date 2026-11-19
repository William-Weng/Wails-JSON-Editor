package main

import (
	"embed"
	"wails-json-editor/backend"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {}

func main() {

	updateService := &backend.UpdateService{}

	app := application.New(application.Options{
		Name:        "JSON Editor",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(updateService),
		},
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

	menu := app.Menu.New()
	menu.AddRole(application.AppMenu)
	app.Menu.Set(menu)

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
