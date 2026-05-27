package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var flagMain string
var Ctx context.Context

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Наша функция для расчетов
func (a *App) Start(flag string, addr string) {
	Ctx = a.ctx

	if flag == "--client" {
		InitializeClient()
		ClientRun(addr)

		flagMain = flag

		WiriteArea(fmt.Sprintf("run - [%s] addres - [%s]", flag, addr), "logMenu")
	} else if flag == "--server" {
		InitializeServer()
		ServerRun(addr)

		flagMain = flag
		WiriteArea(fmt.Sprintf("run - [%s] addres - [%s]", flag, addr), "logMenu")
	}

}

func WiriteArea(msg string, area string) {
	runtime.EventsEmit(Ctx, "writeArea", msg, area)
}

func (a *App) Send(msg string) {
	if flagMain == "--client" {
		WriteClient(msg)
	} else if flagMain == "--server" {
		WriteServer(msg)
	}
}
