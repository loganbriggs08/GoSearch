package main

import (
	"context"
	"fmt"

	"github.com/NotKatsu/GoSearch/keystroke"
)

type App struct {
	ctx context.Context
}

func GoSearch() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go keystroke.Listener(a.ctx)
}

func (a *App) Search(query string) {
	fmt.Println(query)
}
