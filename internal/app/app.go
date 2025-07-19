package app

import (
	"KKonverter/internal/converter"
	"KKonverter/internal/ui"
)


type App struct {
	ui         *ui.AppUI
	converters map[string]converter.Converter
}

func NewApp() *App {
	converters := map[string]converter.Converter{
		".docx": &converter.DocumentConverter{},
		".pdf":  &converter.DocumentConverter{},
		".mp4":  &converter.MediaConverter{},
		".mp3":  &converter.MediaConverter{},
	}

	return &App{
		ui:         ui.NewAppUI(),
		converters: converters,
	}
}

func (a *App) Run() {
	a.ui.LoadUI(a.converters)
	a.ui.Run()
}