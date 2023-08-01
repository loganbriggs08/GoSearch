package main

import (
	"context"
	"fmt"
)
type App struct {
	ctx context.Context
}
func GoSearch() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}


func (a *App) Search(name string) {
	fmt.Printf("Search Result: %s\n", name)
}
