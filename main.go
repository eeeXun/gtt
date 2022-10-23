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
	langButton      = tview.NewButton("(1)Language")
	styleButton     = tview.NewButton("(2)Style")
	menuButton      = tview.NewButton("(3)KeyMap")
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

	mainPage.AddPage("translatePage", translateWindow, true, true)
	mainPage.AddPage("langPage", langWindow, true, false)

	if err := app.SetRoot(mainPage, true).
		EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	defer updateConfig()
}
