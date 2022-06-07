package main

import (
	"embed"

	// hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed frontend/dist
var assets embed.FS

// func shortcuts(fn func()) {
// 	hook.Register(hook.KeyDown, []string{"alt", "t"}, func(e hook.Event) {
// 		fn()
// 	})
// 	s := hook.Start()
// 	hook.Process(s)
// }
func addEmptyMenu(m *menu.Menu, t string, k *keys.Accelerator) {
	m.AddText(t, k, func(_ *menu.CallbackData) {})
}

func main() {
	// low()
	// Create an instance of the app structure
	app := NewApp()
	AppMenu := menu.NewMenu()
	// 此处还需要针对MACOS进行判断
	AppMenu.Append(menu.EditMenu())
	WindowMenu := AppMenu.AddSubmenu("窗口")
	WindowMenu.AddText("最小化", keys.CmdOrCtrl("w"), func(_ *menu.CallbackData) {
		runtime.WindowMinimise(app.ctx)
	})
	WindowMenu.AddText("重载", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		app.Reload()
	})
	addEmptyMenu(WindowMenu, "配置", keys.CmdOrCtrl(","))
	WindowMenu.AddSeparator()
	WindowMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	EditorMenu := AppMenu.AddSubmenu("编辑")
	addEmptyMenu(EditorMenu, "搜索文本处理器", keys.CmdOrCtrl("enter"))
	addEmptyMenu(EditorMenu, "搜索文本处理器", keys.CmdOrCtrl("f"))
	addEmptyMenu(EditorMenu, "清空文本框", keys.CmdOrCtrl("d"))
	addEmptyMenu(EditorMenu, "还原文本框", keys.CmdOrCtrl("z"))
	addEmptyMenu(EditorMenu, "复制文本框内容", keys.OptionOrAlt("c"))
	addEmptyMenu(EditorMenu, "上一版本", keys.CmdOrCtrl("1"))
	addEmptyMenu(EditorMenu, "下一版本", keys.CmdOrCtrl("2"))
	// go shortcuts(func() {
	// 	runtime.WindowShow(app.ctx)
	// })

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Texz",
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
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: true,
			WebviewUserDataPath:               "",
			Theme:                             windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Texz",
				Message: "© 2022 Priate",
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
