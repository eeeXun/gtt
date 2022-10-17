package main

import (
	"fmt"
	"github.com/rivo/tview"
)

var (
	// Translate
	translator Translator
	// TUI
	app      = tview.NewApplication()
	src_box  = tview.NewInputField()
	dest_box = tview.NewTextView()
)

func main() {
	translator.src_lang = "English"
	translator.dest_lang = "Chinese (Traditional)"
	result, _ := translator.Translate("Hello world\nApple\nbumper")
	fmt.Println(result)
	// if err := app.SetRoot(
	// 	tview.NewFlex().SetDirection(tview.FlexRow).
	// 		AddItem(src_box, 0, 1, true).
	// 		AddItem(dest_box, 0, 6, false),
	// 	true).EnableMouse(true).Run(); err != nil {
	// 	panic(err)
	// }
}
