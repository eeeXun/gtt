package main

import (
	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

var (
	// Translate
	translator = NewTranslator()
	// UI
	app             = tview.NewApplication()
	srcBox          = tview.NewTextArea()
	dstBox          = tview.NewTextView()
	srcDropDown     = tview.NewDropDown()
	dstDropDown     = tview.NewDropDown()
	langButton      = tview.NewButton("(1)Lang")
	styleButton     = tview.NewButton("(2)Style")
	menuButton      = tview.NewButton("(3)Menu")
	translateWindow = tview.NewFlex()
	langWindow      = tview.NewFlex()
	mainPage        = tview.NewPages()
	window          Window
	// config
	config      = viper.New()
	theme       string
	transparent bool
)

func main() {
	SetTermTitle("GTT")
	configInit()
	window.colorInit()
	uiInit()
	translateWindow.SetDirection(tview.FlexColumn).
		AddItem(srcBox, 0, 1, true).
		AddItem(dstBox, 0, 1, false)
	langWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(srcDropDown, 32, 1, true).
			AddItem(dstDropDown, 32, 1, false).
			AddItem(nil, 0, 1, false), 20, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(langButton, 7, 1, true).
			AddItem(nil, 20, 1, false).
			AddItem(styleButton, 8, 1, true).
			AddItem(nil, 20, 1, false).
			AddItem(menuButton, 7, 1, true).
			AddItem(nil, 0, 1, false), 1, 1, true).
		AddItem(nil, 0, 1, false)

	mainPage.AddPage("translatePage", translateWindow, true, true)
	mainPage.AddPage("langPage", langWindow, true, false)

	if err := app.SetRoot(mainPage, true).
		EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	defer updateConfig()
}
