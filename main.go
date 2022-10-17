package main

import (
	// "fmt"
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	// Translate
	translator Translator
	// UI
	app      = tview.NewApplication()
	src_box  = tview.NewTextArea()
	dest_box = tview.NewTextView()
	window   Window
)

func main() {
	translator.src_lang = "English"
	translator.dest_lang = "Chinese (Traditional)"
	// result, _ := translator.Translate("Hello world\nApple\nbumper")
	// fmt.Println(result)
	window.color_init()
	ui_init()
	if err := app.SetRoot(
		tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(src_box, 0, 1, true).
			AddItem(dest_box, 0, 1, false),
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
