package main

import (
	"context"
	"texz/config"
	"texz/str"

	"github.com/atotto/clipboard"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetModules() []str.Module {
	return str.GetModules()
}

func (a *App) Handling(funcName string, rawText string) string {
	f := str.GetFunction(funcName)
	if f == nil {
		return "处理失败，模块不存在!"
	}
	return f.Call(rawText)
}
func (a *App) Copy(text string) error {
	res := clipboard.WriteAll(text)
	return res
}

func (a *App) Reload() {
	runtime.WindowReload(a.ctx)
	runtime.WindowReloadApp(a.ctx)
}

func (a *App) GetConfig() string {
	return config.GetConfig()
}

func (a *App) GetRemoteModlues() []string {
	return config.GetRemoteModlues()
}
