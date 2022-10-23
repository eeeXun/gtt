package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func updateBackground() {
	// box
	srcBox.SetBackgroundColor(window.src.backgroundColor)
	srcBox.SetTextStyle(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor))
	dstBox.SetBackgroundColor(window.dst.backgroundColor)

	// dropdown
	srcLangDropDown.SetBackgroundColor(window.src.backgroundColor)
	srcLangDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor),
		tcell.StyleDefault.
			Background(window.src.selectedColor).
			Foreground(window.src.prefixColor))
	dstLangDropDown.SetBackgroundColor(window.dst.backgroundColor)
	dstLangDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor),
		tcell.StyleDefault.
			Background(window.src.selectedColor).
			Foreground(window.src.prefixColor))
	themeDropDown.SetBackgroundColor(window.src.backgroundColor)
	themeDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor),
		tcell.StyleDefault.
			Background(window.src.selectedColor).
			Foreground(window.src.prefixColor))
	transparentDropDown.SetBackgroundColor(window.src.backgroundColor)
	transparentDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor),
		tcell.StyleDefault.
			Background(window.src.selectedColor).
			Foreground(window.src.prefixColor))
}

// update title and option
func updateTitle() {
	srcBox.SetTitle(translator.srcLang)
	dstBox.SetTitle(translator.dstLang)
	srcLangDropDown.SetCurrentOption(IndexOf(translator.srcLang, Lang))
	srcLangDropDown.SetTitle(translator.srcLang)
	dstLangDropDown.SetCurrentOption(IndexOf(translator.dstLang, Lang))
	dstLangDropDown.SetTitle(translator.dstLang)
}

func attachButton() *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(langButton, 11, 1, true).
		AddItem(nil, 18, 1, false).
		AddItem(styleButton, 8, 1, true).
		AddItem(nil, 18, 1, false).
		AddItem(menuButton, 9, 1, true).
		AddItem(nil, 0, 1, false)
}

func uiInit() {
	// box
	srcBox.SetBorder(true).
		SetBorderColor(window.src.borderColor).
		SetTitleColor(window.src.borderColor)
	srcBox.SetSelectedStyle(tcell.StyleDefault.
		Background(window.src.selectedColor).
		Foreground(window.src.foregroundColor))
	dstBox.SetBorder(true).
		SetBorderColor(window.dst.borderColor).
		SetTitleColor(window.dst.borderColor)
	dstBox.SetTextColor(window.dst.foregroundColor)

	// dropdown
	srcLangDropDown.SetOptions(Lang, nil)
	srcLangDropDown.SetFieldBackgroundColor(window.src.selectedColor).
		SetFieldTextColor(window.src.foregroundColor).
		SetPrefixTextColor(window.src.prefixColor)
	srcLangDropDown.SetBorder(true).
		SetBorderColor(window.src.borderColor).
		SetTitleColor(window.src.borderColor)
	dstLangDropDown.SetOptions(Lang, nil)
	dstLangDropDown.SetFieldBackgroundColor(window.src.selectedColor).
		SetFieldTextColor(window.src.foregroundColor).
		SetPrefixTextColor(window.src.prefixColor)
	dstLangDropDown.SetBorder(true).
		SetBorderColor(window.dst.borderColor).
		SetTitleColor(window.dst.borderColor)
	themeDropDown.SetOptions(themes_name, nil).
		SetLabel("Theme: ").SetLabelColor(window.src.labelColor)
	themeDropDown.SetFieldBackgroundColor(window.src.selectedColor).
		SetFieldTextColor(window.src.foregroundColor).
		SetPrefixTextColor(window.src.prefixColor)
	transparentDropDown.SetOptions([]string{"true", "false"}, nil).
		SetLabel("Transparent: ").SetLabelColor(window.src.labelColor)
	transparentDropDown.SetFieldBackgroundColor(window.src.selectedColor).
		SetFieldTextColor(window.src.foregroundColor).
		SetPrefixTextColor(window.src.prefixColor)

	// button
	langButton.SetLabelColor(window.src.foregroundColor).
		SetBackgroundColorActivated(window.src.pressColor).
		SetLabelColorActivated(window.src.foregroundColor).
		SetBackgroundColor(window.src.selectedColor)
	styleButton.SetLabelColor(window.src.foregroundColor).
		SetBackgroundColorActivated(window.src.pressColor).
		SetLabelColorActivated(window.src.foregroundColor).
		SetBackgroundColor(window.src.selectedColor)
	menuButton.SetLabelColor(window.src.foregroundColor).
		SetBackgroundColorActivated(window.src.pressColor).
		SetLabelColorActivated(window.src.foregroundColor).
		SetBackgroundColor(window.src.selectedColor)

	updateBackground()
	updateTitle()

	// window
	translateWindow.SetDirection(tview.FlexColumn).
		AddItem(srcBox, 0, 1, true).
		AddItem(dstBox, 0, 1, false)
	langWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(srcLangDropDown, 32, 1, true).
			AddItem(dstLangDropDown, 32, 1, false).
			AddItem(nil, 0, 1, false), 20, 1, true).
		AddItem(attachButton(), 1, 1, true).
		AddItem(nil, 0, 1, false)
	styleWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(themeDropDown, 1, 1, false).
				AddItem(transparentDropDown, 1, 1, false), 20, 1, false).
			AddItem(nil, 0, 1, false), 20, 1, true).
		AddItem(attachButton(), 1, 1, true).
		AddItem(nil, 0, 1, false)

	// handler
	mainPage.SetInputCapture(pagesHandler)
	translateWindow.SetInputCapture(translatePageHandler)
	srcLangDropDown.SetDoneFunc(srcDropDownHandler).
		SetSelectedFunc(srcSelected)
	dstLangDropDown.SetDoneFunc(dstDropDownHandler).
		SetSelectedFunc(dstSelected)
}

func pagesHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyCtrlT:
		if transparent {
			window.src.backgroundColor = Themes[theme]["bg"]
			window.dst.backgroundColor = Themes[theme]["bg"]
		} else {
			window.src.backgroundColor = Transparent
			window.dst.backgroundColor = Transparent
		}
		updateBackground()
		transparent = !transparent
	}

	return event
}

func translatePageHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyEsc:
		mainPage.ShowPage("langPage")
	case tcell.KeyCtrlJ:
		message := srcBox.GetText()
		if len(message) > 0 {
			result, err := translator.Translate(message)
			if err != nil {
				dstBox.SetText(err.Error())
			} else {
				dstBox.SetText(result)
			}
		}
	case tcell.KeyCtrlQ:
		srcBox.SetText("", true)
	case tcell.KeyCtrlS:
		translator.srcLang, translator.dstLang = translator.dstLang, translator.srcLang
		updateTitle()
		src_text := srcBox.GetText()
		dst_text := dstBox.GetText(false)
		if len(dst_text) > 0 {
			// GetText of Box contains "\n" if it has words
			srcBox.SetText(dst_text[:len(dst_text)-1], true)
		} else {
			srcBox.SetText(dst_text, true)
		}
		dstBox.SetText(src_text)
	case tcell.KeyCtrlO:
		// play source sound
		if translator.soundLock.Available() {
			message := srcBox.GetText()
			if len(message) > 0 {
				translator.soundLock.Acquire()
				go func() {
					err := translator.PlaySound(translator.srcLang, message)
					if err != nil {
						srcBox.SetText(err.Error(), true)
					}
				}()
			}

		}
	case tcell.KeyCtrlP:
		// play destination sound
		if translator.soundLock.Available() {
			message := dstBox.GetText(false)
			if len(message) > 0 {
				translator.soundLock.Acquire()
				go func() {
					err := translator.PlaySound(translator.dstLang, message)
					if err != nil {
						dstBox.SetText(err.Error())
					}
				}()
			}
		}
	case tcell.KeyCtrlX:
		// stop play sound
		translator.soundLock.stop = true
	}

	return event
}

func srcSelected(text string, index int) {
	translator.srcLang = text
	srcBox.SetTitle(text)
	srcLangDropDown.SetTitle(text)
}

func dstSelected(text string, index int) {
	translator.dstLang = text
	dstBox.SetTitle(text)
	dstLangDropDown.SetTitle(text)
}

func srcDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(dstLangDropDown)
	case tcell.KeyEsc:
		mainPage.HidePage("langPage")
	}
}

func dstDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(srcLangDropDown)
	case tcell.KeyEsc:
		mainPage.HidePage("langPage")
	}
}
