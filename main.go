package main

import (
	"flag"
	"gtt/internal/translate"
	"gtt/internal/ui"

	"github.com/rivo/tview"
)

var (
	// version
	version string
	// argument
	srcLangArg *string = flag.String("src", "", "Set source language")
	dstLangArg *string = flag.String("dst", "", "Set destination language")
	// Translate
	translator  translate.Translator
	translators = map[string]translate.Translator{
		"ArgosTranslate":  translate.NewLibreTranslate(),
		"GoogleTranslate": translate.NewGoogleTranslate(),
	}
	// UI
	app                 = tview.NewApplication()
	srcInput            = tview.NewTextArea()
	dstOutput           = tview.NewTextView()
	defOutput           = tview.NewTextArea()
	posOutput           = tview.NewTextArea()
	translatorDropDown  = tview.NewDropDown()
	srcLangDropDown     = tview.NewDropDown()
	dstLangDropDown     = tview.NewDropDown()
	langCycle           = ui.NewUICycle(translatorDropDown, srcLangDropDown, dstLangDropDown)
	themeDropDown       = tview.NewDropDown()
	transparentDropDown = tview.NewDropDown()
	hideBelowDropDown   = tview.NewDropDown()
	srcBorderDropDown   = tview.NewDropDown()
	dstBorderDropDown   = tview.NewDropDown()
	styleCycle          = ui.NewUICycle(
		themeDropDown,
		transparentDropDown,
		hideBelowDropDown,
		srcBorderDropDown,
		dstBorderDropDown)
	keyMapMenu           = tview.NewTextView()
	langButton           = tview.NewButton("(1)Language")
	styleButton          = tview.NewButton("(2)Style")
	keyMapButton         = tview.NewButton("(3)KeyMap")
	translateWindow      = tview.NewFlex()
	translateAboveWidget = tview.NewFlex()
	translateBelowWidget = tview.NewFlex()
	langWindow           = tview.NewFlex()
	styleWindow          = tview.NewFlex()
	keyMapWindow         = tview.NewFlex()
	mainPage             = tview.NewPages()
)

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	switch {
	case *showVersion:
		print(version, "\n")
	default:
		SetTermTitle("GTT")
		configInit()
		uiInit()

		mainPage.AddPage("translateWindow", translateWindow, true, true)
		mainPage.AddPage("langWindow", langWindow, true, false)
		mainPage.AddPage("styleWindow", styleWindow, true, false)
		mainPage.AddPage("keyMapWindow", keyMapWindow, true, false)

		if err := app.SetRoot(mainPage, true).
			EnableMouse(true).Run(); err != nil {
			panic(err)
		}

		// Check if config need to update
		defer updateConfig()
	}
}
