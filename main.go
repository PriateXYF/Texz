package main

import (
	"embed"

	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed frontend/dist
var assets embed.FS

func shortcuts(fn func()) {
	hook.Register(hook.KeyDown, []string{"alt", "t"}, func(e hook.Event) {
		fn()
	})
	s := hook.Start()
	hook.Process(s)
}
func main() {
	// low()
	// Create an instance of the app structure
	app := NewApp()
	AppMenu := menu.NewMenu()
	// 此处还需要针对MACOS进行判断
	AppMenu.Append(menu.EditMenu())
	WindowMenu := AppMenu.AddSubmenu("Window")
	WindowMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
	WindowMenu.AddSeparator()
	WindowMenu.AddText("Close", keys.CmdOrCtrl("w"), func(_ *menu.CallbackData) {
		runtime.WindowMinimise(app.ctx)
	})
	WindowMenu.AddSeparator()
	WindowMenu.AddText("Reload", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		runtime.WindowReload(app.ctx)
	})
	go shortcuts(func() {
		runtime.WindowShow(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Text Tool",
		Width:         432,
		Height:        312,
		MaxHeight:     768,
		DisableResize: true,
		// AlwaysOnTop:       true,
		HideWindowOnClose: true,
		Assets:            assets,
		OnStartup:         app.startup,
		Menu:              AppMenu,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Text Tool",
				Message: "© 2022 Priate",
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
