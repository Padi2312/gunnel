package main

import (
	"context"
	"embed"
	"fmt"
	"gunnel/internal"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const version = "0.1.0"

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	store := internal.NewStore()
	store.Init()

	handler := internal.NewHandler()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  fmt.Sprintf("Gunnel %s", version),
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			store.Startup(ctx)
			handler.Startup(ctx)
		},
		Bind: []interface{}{
			store,
			handler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
