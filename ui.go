package main

import (
	"github.com/gdamore/tcell/v2"
)

func updateBackground() {
	// box
	srcBox.SetBackgroundColor(window.src.backgroundColor)
	srcBox.SetTextStyle(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor))

	dstBox.SetBackgroundColor(window.dst.backgroundColor)

	// dropdown
	srcDropDown.SetBackgroundColor(window.src.backgroundColor)
	srcDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.backgroundColor).
		Foreground(window.src.foregroundColor),
		tcell.StyleDefault.
			Background(window.src.selectedColor).
			Foreground(window.src.prefixColor))

	dstDropDown.SetBackgroundColor(window.dst.backgroundColor)
	dstDropDown.SetListStyles(tcell.StyleDefault.
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
	srcDropDown.SetCurrentOption(IndexOf(translator.srcLang, Lang))
	srcDropDown.SetTitle(translator.srcLang)
	dstDropDown.SetCurrentOption(IndexOf(translator.dstLang, Lang))
	dstDropDown.SetTitle(translator.dstLang)
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
	srcDropDown.SetOptions(Lang, nil)
	srcDropDown.SetFieldBackgroundColor(window.src.selectedColor).
		SetFieldTextColor(window.src.foregroundColor).
		SetPrefixTextColor(window.dst.prefixColor)
	srcDropDown.SetBorder(true).
		SetBorderColor(window.src.borderColor).
		SetTitleColor(window.src.borderColor)

	dstDropDown.SetOptions(Lang, nil)
	dstDropDown.SetFieldBackgroundColor(window.src.selectedColor).
		SetFieldTextColor(window.src.foregroundColor).
		SetPrefixTextColor(window.dst.prefixColor)
	dstDropDown.SetBorder(true).
		SetBorderColor(window.dst.borderColor).
		SetTitleColor(window.dst.borderColor)

	updateBackground()
	updateTitle()

	// handler
	mainPage.SetInputCapture(pagesHandler)
	translateWindow.SetInputCapture(translatePageHandler)
	srcDropDown.SetDoneFunc(srcDropDownHandler).
		SetSelectedFunc(srcSelected)
	dstDropDown.SetDoneFunc(dstDropDownHandler).
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
	srcDropDown.SetTitle(text)
}

func dstSelected(text string, index int) {
	translator.dstLang = text
	dstBox.SetTitle(text)
	dstDropDown.SetTitle(text)
}

func srcDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(dstDropDown)
	case tcell.KeyEsc:
		mainPage.HidePage("langPage")
	}
}

func dstDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(srcDropDown)
	case tcell.KeyEsc:
		mainPage.HidePage("langPage")
	}
}
