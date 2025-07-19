package ui

import (
	"KKonverter/internal/converter"
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type AppUI struct {
	app        fyne.App
	window     fyne.Window
	converters map[string]converter.Converter

	fileList           []string
	fileListWidget     *widget.List
	outputFormatSelect *widget.Select
	outputDirEntry     *widget.Entry
	outputDirButton    *widget.Button
	presetSelect       *widget.Select
	pdfEngineSelect    *widget.Select
	progressBar        *widget.ProgressBar
	convertButton      *widget.Button
}

func NewAppUI() (*AppUI, string) {
	a := app.NewWithID("com.kkonverter.app")
	w := a.NewWindow("KKonverter")
	w.SetMaster()

	ui := &AppUI{
		app:    a,
		window: w,
	}

	w.SetOnDropped(func(p fyne.Position, uris []fyne.URI) {
		for _, u := range uris {
			ui.fileList = append(ui.fileList, u.Path())
		}
		ui.fileListWidget.Refresh()
	})

	ui.fileList = []string{}
	ui.fileListWidget = widget.NewList(
		func() int {
			return len(ui.fileList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ui.fileList[i])
		},
	)
	ui.outputFormatSelect = widget.NewSelect([]string{".docx", ".pdf", ".mp4", ".mp3"}, nil)
	ui.outputDirEntry = widget.NewEntry()
	ui.outputDirEntry.SetText(".") // Default to current directory
	ui.outputDirButton = widget.NewButton("Browse...", ui.selectOutputDir)

	presets := []string{"None", "High Quality", "Medium Quality", "Low Quality"}
	ui.presetSelect = widget.NewSelect(presets, func(selected string) {
		switch selected {
		case "High Quality":
		case "Medium Quality":
		case "Low Quality":
		}
	})
	ui.presetSelect.SetSelected("None")

	ui.progressBar = widget.NewProgressBar()
	ui.progressBar.Hide()

	pdfEngines := []string{"default", "pdflatex", "xelatex", "lualatex", "wkhtmltopdf", "weasyprint", "prince"}
	ui.pdfEngineSelect = widget.NewSelect(pdfEngines, nil)
	ui.pdfEngineSelect.SetSelected("default")

	ui.convertButton = widget.NewButton("Convert All", ui.convertFiles)

	return ui, ui.pdfEngineSelect.Selected
}

func (a *AppUI) LoadUI(converters map[string]converter.Converter) {
	a.converters = converters
	addFileButton := widget.NewButton("Add Files", a.addFiles)

	outputDirContainer := container.NewBorder(
		nil, nil,
		widget.NewLabel("Output Directory:"),
		a.outputDirButton,
		a.outputDirEntry,
	)

	content := container.NewVBox(
		addFileButton,
		a.fileListWidget,
		widget.NewLabel("Output Format:"),
		a.outputFormatSelect,
		widget.NewLabel("Preset:"),
		a.presetSelect,
		widget.NewLabel("PDF Engine:"),
		a.pdfEngineSelect,
		outputDirContainer,
		a.progressBar,
		a.convertButton,
	)

	a.window.SetContent(content)
}

func (a *AppUI) selectOutputDir() {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if uri == nil {
			return
		}
		a.outputDirEntry.SetText(uri.Path())
	}, a.window)
}

func (a *AppUI) addFiles() {
	dialog.ShowFileOpen(func(files fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if files == nil {
			return
		}
		a.fileList = append(a.fileList, files.URI().Path())
		a.fileListWidget.Refresh()
	}, a.window)
}

func (a *AppUI) convertFiles() {
	if len(a.fileList) == 0 {
		dialog.ShowError(fmt.Errorf("no input files selected"), a.window)
		return
	}

	outputFormat := a.outputFormatSelect.Selected
	if outputFormat == "" {
		dialog.ShowError(fmt.Errorf("no output format selected"), a.window)
		return
	}

	outputDir := a.outputDirEntry.Text
	if outputDir == "" {
		outputDir = "."
	}

	a.progressBar.Show()
	a.progressBar.SetValue(0)

	totalFiles := float64(len(a.fileList))
	for i, inputFile := range a.fileList {
		baseName := filepath.Base(inputFile[0 : len(inputFile)-len(filepath.Ext(inputFile))])
		outputFile := filepath.Join(outputDir, baseName+outputFormat)

		converter, ok := a.converters[filepath.Ext(inputFile)]
		if !ok {
			dialog.ShowError(fmt.Errorf("no converter found for %s files", filepath.Ext(inputFile)), a.window)
			continue
		}

		err := converter.Convert(inputFile, outputFile)
		if err != nil {
			dialog.ShowError(err, a.window)
			continue
		}
		a.progressBar.SetValue(float64(i+1) / totalFiles)
	}

	dialog.ShowInformation("Success", "All files converted successfully", a.window)
	a.fileList = []string{}
	a.fileListWidget.Refresh()
	a.progressBar.Hide()
}

func (a *AppUI) Run() {
	a.window.ShowAndRun()
}
