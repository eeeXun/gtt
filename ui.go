package main

import (
	"github.com/gdamore/tcell/v2"
)

func updateBackground() {
	// box
	src_box.SetBackgroundColor(window.src.background_color)
	src_box.SetTextStyle(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color))

	dst_box.SetBackgroundColor(window.dst.background_color)

	// dropdown
	src_dropdown.SetBackgroundColor(window.src.background_color)
	src_dropdown.SetListStyles(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color),
		tcell.StyleDefault.
			Background(window.src.selected_color).
			Foreground(window.src.prefix_color))

	dst_dropdown.SetBackgroundColor(window.dst.background_color)
	dst_dropdown.SetListStyles(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color),
		tcell.StyleDefault.
			Background(window.src.selected_color).
			Foreground(window.src.prefix_color))
}

// update title and option
func updateTitle() {
	src_box.SetTitle(translator.srcLang)
	dst_box.SetTitle(translator.dstLang)
	src_dropdown.SetCurrentOption(IndexOf(translator.srcLang, Lang))
	src_dropdown.SetTitle(translator.srcLang)
	dst_dropdown.SetCurrentOption(IndexOf(translator.dstLang, Lang))
	dst_dropdown.SetTitle(translator.dstLang)
}

func uiInit() {
	// box
	src_box.SetBorder(true).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color)
	src_box.SetSelectedStyle(tcell.StyleDefault.
		Background(window.src.selected_color).
		Foreground(window.src.foreground_color))

	dst_box.SetBorder(true).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.border_color)
	dst_box.SetTextColor(window.dst.foreground_color)

	// dropdown
	src_dropdown.SetOptions(Lang, nil)
	src_dropdown.SetFieldBackgroundColor(window.src.selected_color).
		SetFieldTextColor(window.src.foreground_color).
		SetPrefixTextColor(window.dst.prefix_color)
	src_dropdown.SetBorder(true).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color)

	dst_dropdown.SetOptions(Lang, nil)
	dst_dropdown.SetFieldBackgroundColor(window.src.selected_color).
		SetFieldTextColor(window.src.foreground_color).
		SetPrefixTextColor(window.dst.prefix_color)
	dst_dropdown.SetBorder(true).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.border_color)

	updateBackground()
	updateTitle()

	// handler
	pages.SetInputCapture(pagesHandler)
	translate_page.SetInputCapture(translatePageHandler)
	src_dropdown.SetDoneFunc(srcDropDownHandler).
		SetSelectedFunc(srcSelected)
	dst_dropdown.SetDoneFunc(dstDropDownHandler).
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
		pages.ShowPage("lang_page")
	case tcell.KeyCtrlJ:
		message := src_box.GetText()
		if len(message) > 0 {
			result, err := translator.Translate(message)
			if err != nil {
				dst_box.SetText(err.Error())
			} else {
				dst_box.SetText(result)
			}
		}
	case tcell.KeyCtrlQ:
		src_box.SetText("", true)
	case tcell.KeyCtrlS:
		translator.srcLang, translator.dstLang = translator.dstLang, translator.srcLang
		updateTitle()
		src_text := src_box.GetText()
		dst_text := dst_box.GetText(false)
		if len(dst_text) > 0 {
			// GetText of Box contains "\n" if it has words
			src_box.SetText(dst_text[:len(dst_text)-1], true)
		} else {
			src_box.SetText(dst_text, true)
		}
		dst_box.SetText(src_text)
	case tcell.KeyCtrlO:
		// play source sound
		if translator.soundLock.Available() {
			message := src_box.GetText()
			if len(message) > 0 {
				translator.soundLock.Acquire()
				go func() {
					err := translator.PlaySound(translator.srcLang, message)
					if err != nil {
						src_box.SetText(err.Error(), true)
					}
				}()
			}

		}
	case tcell.KeyCtrlP:
		// play destination sound
		if translator.soundLock.Available() {
			message := dst_box.GetText(false)
			if len(message) > 0 {
				translator.soundLock.Acquire()
				go func() {
					err := translator.PlaySound(translator.dstLang, message)
					if err != nil {
						dst_box.SetText(err.Error())
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
	src_box.SetTitle(text)
	src_dropdown.SetTitle(text)
}

func dstSelected(text string, index int) {
	translator.dstLang = text
	dst_box.SetTitle(text)
	dst_dropdown.SetTitle(text)
}

func srcDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(dst_dropdown)
	case tcell.KeyEsc:
		pages.HidePage("lang_page")
	}
}

func dstDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(src_dropdown)
	case tcell.KeyEsc:
		pages.HidePage("lang_page")
	}
}
