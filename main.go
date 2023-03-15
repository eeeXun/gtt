package main

import (
	"flag"

	"github.com/eeeXun/gtt/internal/style"
	"github.com/eeeXun/gtt/internal/translate"
	"github.com/eeeXun/gtt/internal/ui"
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
	translators = make(map[string]translate.Translator, len(translate.AllTranslator))
	// UI style
	uiStyle = style.NewStyle()
	// UI
	app                 = tview.NewApplication()
	srcInput            = tview.NewTextArea()
	dstOutput           = tview.NewTextView()
	defOutput           = tview.NewTextArea()
	posOutput           = tview.NewTextArea()
	translatorDropDown  = tview.NewDropDown()
	srcLangDropDown     = tview.NewDropDown()
	dstLangDropDown     = tview.NewDropDown()
	langCycle           = ui.NewUICycle(srcLangDropDown, dstLangDropDown, translatorDropDown)
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
	langPopOut           = tview.NewFlex()
	stylePopOut          = tview.NewFlex()
	keyMapPopOut         = tview.NewFlex()
	mainPage             = tview.NewPages()
)

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	switch {
	case *showVersion:
		print(version, "\n")
	default:
		configInit()
		uiInit()
		SetTermTitle(translator.GetEngineName())

		mainPage.AddPage("translateWindow", translateWindow, true, true)
		mainPage.AddPage("langPopOut", langPopOut, true, false)
		mainPage.AddPage("stylePopOut", stylePopOut, true, false)
		mainPage.AddPage("keyMapPopOut", keyMapPopOut, true, false)

		if err := app.SetRoot(mainPage, true).
			EnableMouse(true).Run(); err != nil {
			panic(err)
		}

		// Check if config file need to be updated
		defer updateConfig()
	}
}
