package main

import (
	"fmt"
	"github.com/rivo/tview"
)

var (
	app        = tview.NewApplication()
	source_box = tview.NewInputField()
	target_box = tview.NewTextView()
)

func main() {
	result := Translate("Hello world\nApple", "English", "Chinese (Traditional)")
	fmt.Println(result)
	// if err := app.SetRoot(
	// 	tview.NewFlex().SetDirection(tview.FlexRow).
	// 		AddItem(source_box, 0, 1, true).
	// 		AddItem(target_box, 0, 6, false),
	// 	true).EnableMouse(true).Run(); err != nil {
	// 	panic(err)
	// }
}
