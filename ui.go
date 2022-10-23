package main

import (
	"github.com/gdamore/tcell/v2"
)

func updateBackground() {
	// box
	srcBox.SetBackgroundColor(window.src.background_color)
	srcBox.SetTextStyle(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color))

	dstBox.SetBackgroundColor(window.dst.background_color)

	// dropdown
	srcDropDown.SetBackgroundColor(window.src.background_color)
	srcDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color),
		tcell.StyleDefault.
			Background(window.src.selected_color).
			Foreground(window.src.prefix_color))

	dstDropDown.SetBackgroundColor(window.dst.background_color)
	dstDropDown.SetListStyles(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color),
		tcell.StyleDefault.
			Background(window.src.selected_color).
			Foreground(window.src.prefix_color))
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
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color)
	srcBox.SetSelectedStyle(tcell.StyleDefault.
		Background(window.src.selected_color).
		Foreground(window.src.foreground_color))

	dstBox.SetBorder(true).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.border_color)
	dstBox.SetTextColor(window.dst.foreground_color)

	// dropdown
	srcDropDown.SetOptions(Lang, nil)
	srcDropDown.SetFieldBackgroundColor(window.src.selected_color).
		SetFieldTextColor(window.src.foreground_color).
		SetPrefixTextColor(window.dst.prefix_color)
	srcDropDown.SetBorder(true).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color)

	dstDropDown.SetOptions(Lang, nil)
	dstDropDown.SetFieldBackgroundColor(window.src.selected_color).
		SetFieldTextColor(window.src.foreground_color).
		SetPrefixTextColor(window.dst.prefix_color)
	dstDropDown.SetBorder(true).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.border_color)

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
			window.src.background_color = Themes[theme]["bg"]
			window.dst.background_color = Themes[theme]["bg"]
		} else {
			window.src.background_color = Transparent
			window.dst.background_color = Transparent
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
