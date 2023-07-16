package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	Menu *Menu
	Ui   fyne.App
}

var a *App

func init() {
	a = &App{
		Menu: CreateMenu(),
		Ui:   app.New(),
	}
}

func Start() error {
	a.Menu.Show(a.Ui)

	return nil
}
