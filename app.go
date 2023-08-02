package main

import (
	"context"
	"fmt"
	"github.com/NotKatsu/GoSearch/modules"
)

type App struct {
	ctx context.Context
}
func GoSearch() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go modules.KeystrokeListener(a.ctx)
}


func (a *App) Search(name string) {
	fmt.Printf("Search Result: %s\n", name)
}
