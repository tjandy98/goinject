package main

import (
	"fmt"
	"image/color"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func gui() {
	processList, _ := getProcesses()

	a := app.New()
	w := a.NewWindow("GoInject")

	var selectedProcess = widget.NewLabel("-")
	var selectedProcessId = widget.NewLabel("-")

	processesList := widget.NewList(

		func() int {
			return len(processList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(processList[i].Exe)
		})

	processesList.OnSelected = func(id widget.ListItemID) {

		selectedProcess.SetText(processList[id].Exe)
		selectedProcessId.SetText(fmt.Sprintf("%v", processList[id].ProcessID))

	}

	selectProcessLabel := widget.NewLabel("Select Process:")

	processListPane := container.New(layout.NewMaxLayout(), processesList)
	processRefreshButton := widget.NewButton("Reload", func() { processList, _ = getProcesses() })
	processPane := container.New(layout.NewBorderLayout(selectProcessLabel, processRefreshButton, nil, nil), selectProcessLabel, processRefreshButton, processListPane)

	dllInput := widget.NewEntry()
	dllInput.SetPlaceHolder("C:\\go.dll")
	dllInput.OnChanged = func(s string) { fmt.Println(s) }
	browseDllLabel := widget.NewLabel("DLL Path: ")
	browseDllButton := widget.NewButton("Browse", func() { showFilePicker(w, dllInput) })
	selectDllPane := container.New(layout.NewFormLayout(), browseDllLabel, dllInput)

	browsePane := container.New(layout.NewBorderLayout(selectDllPane, nil, nil, nil), selectDllPane, browseDllButton)

	injectionMethodLabel := widget.NewLabel("Injection Method:")
	methodSelect := widget.NewSelect([]string{"LoadLibraryA"}, func(value string) {
	})
	methodSelect.Selected = methodSelect.Options[0]
	injectionMethodPane := container.New(layout.NewFormLayout(), injectionMethodLabel, methodSelect)

	processLabel := widget.NewLabel("Target Process:     ")
	injectionProcessPane := container.New(layout.NewFormLayout(), processLabel, selectedProcess)

	processIdLabel := widget.NewLabel("Target Process ID:")
	injectionPidPane := container.New(layout.NewFormLayout(), processIdLabel, selectedProcessId)

	injectButton := widget.NewButton("Inject", func() {
		i, _ := strconv.Atoi(selectedProcessId.Text)
		injectDll(dllInput.Text, i)
	})

	bottomPane := container.New(layout.NewVBoxLayout(), injectionMethodPane, injectionProcessPane, injectionPidPane, injectButton)

	hSplit := container.NewHSplit(
		processPane,
		browsePane,
	)

	vSplit := container.NewVSplit(hSplit, bottomPane)

	url, _ := url.Parse("https://github.com/tjandy98/goinject")
	text := canvas.NewText("GoInject", color.White)
	arch := widget.NewLabel("64-bit")
	if getProcessArch() {
		arch = widget.NewLabel("32-bit")
	}

	exeInfo := container.NewVBox(
		text,
		arch,
	)
	hyperlink := widget.NewHyperlink("GitHub", url)
	sep := widget.NewSeparator()

	content := container.New(layout.NewVBoxLayout(), exeInfo, sep, hyperlink)

	tabs := container.NewAppTabs(
		container.NewTabItem("Home", vSplit),
		container.NewTabItem("About", container.NewCenter(content)),
	)

	hSplit.SetOffset(0.4)

	w.Resize(fyne.NewSize(700, 400))

	w.SetContent(tabs)

	w.ShowAndRun()

}

func showFilePicker(w fyne.Window, e *widget.Entry) {
	onSelect := func(f fyne.URIReadCloser, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}
		if f == nil {
			return
		}

		e.SetText(f.URI().Path())

	}
	dialog.ShowFileOpen(onSelect, w)
}
