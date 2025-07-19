package app

import (
	"KKonverter/internal/converter"
	"KKonverter/internal/ui"
)

type App struct {
	ui         *ui.AppUI
	converters map[string]converter.Converter
}

func NewApp(pdfEngine string) *App {
	converters := map[string]converter.Converter{
		".docx": &converter.DocumentConverter{PDFEngine: pdfEngine},
		".pdf":  &converter.DocumentConverter{PDFEngine: pdfEngine},
		".mp4":  &converter.MediaConverter{},
		".mp3":  &converter.MediaConverter{},
		".jpg":  &converter.ImageConverter{},
		".jpeg": &converter.ImageConverter{},
		".png":  &converter.ImageConverter{},
	}

	appUI, _ := ui.NewAppUI()
	return &App{
		ui:         appUI,
		converters: converters,
	}
}

func (a *App) Run() {
	a.ui.LoadUI(a.converters)
	a.ui.Run()
}
