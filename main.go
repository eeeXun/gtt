package main

import (
	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

var (
	// Translate
	translator Translator
	// UI
	app            = tview.NewApplication()
	src_box        = tview.NewTextArea()
	dst_box        = tview.NewTextView()
	src_dropdown   = tview.NewDropDown()
	dst_dropdown   = tview.NewDropDown()
	translate_page = tview.NewFlex()
	lang_page      = tview.NewFlex()
	pages          = tview.NewPages()
	window         Window
	// config
	config      = viper.New()
	theme       string
	transparent bool
)

func main() {
	// result, _ := translator.Translate("Hello world\nApple\nbumper")
	// fmt.Println(result)
	configInit()
	window.colorInit()
	uiInit()
	translate_page.SetDirection(tview.FlexColumn).
		AddItem(src_box, 0, 1, true).
		AddItem(dst_box, 0, 1, false)
	lang_page.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(src_dropdown, 32, 1, true).
			AddItem(dst_dropdown, 32, 1, false).
			AddItem(nil, 0, 1, false), 20, 1, true).
		AddItem(nil, 0, 1, false)

	pages.AddPage("translate_page", translate_page, true, true)
	pages.AddPage("lang_page", lang_page, true, false)

	if err := app.SetRoot(pages, true).
		EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	defer updateConfig()
}
