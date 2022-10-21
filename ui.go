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

func updateTitle() {
	src_box.SetTitle(translator.src_lang)
	dst_box.SetTitle(translator.dst_lang)
	src_dropdown.SetCurrentOption(IndexOf(translator.src_lang, Lang))
	src_dropdown.SetTitle(translator.src_lang)
	dst_dropdown.SetCurrentOption(IndexOf(translator.dst_lang, Lang))
	dst_dropdown.SetTitle(translator.dst_lang)
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
	pages.SetInputCapture(PagesHandler)
	translate_page.SetInputCapture(TranslatePageHandler)
	src_dropdown.SetDoneFunc(SrcDropDownHandler).
		SetSelectedFunc(SrcSelected)
	dst_dropdown.SetDoneFunc(DstDropDownHandler).
		SetSelectedFunc(DstSelected)
}

func PagesHandler(event *tcell.EventKey) *tcell.EventKey {
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

func TranslatePageHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyEsc:
		pages.ShowPage("lang_page")
	case tcell.KeyCtrlJ:
		result, err := translator.Translate(src_box.GetText())
		if err != nil {
			dst_box.SetText(err.Error())
		} else {
			dst_box.SetText(result)
		}
	case tcell.KeyCtrlQ:
		src_box.SetText("", true)
	case tcell.KeyCtrlN:
		translator.PlaySound(translator.src_lang, src_box.GetText())
	case tcell.KeyCtrlP:
		translator.PlaySound(translator.dst_lang, dst_box.GetText(false))
	case tcell.KeyCtrlS:
		translator.src_lang, translator.dst_lang = translator.dst_lang, translator.src_lang
		updateTitle()
		src_text := src_box.GetText()
		dst_text := dst_box.GetText(false)
		src_box.SetText(dst_text, true)
		dst_box.SetText(src_text)
	}

	return event
}

func SrcSelected(text string, index int) {
	translator.src_lang = text
	src_box.SetTitle(text)
	src_dropdown.SetTitle(text)
}

func DstSelected(text string, index int) {
	translator.dst_lang = text
	dst_box.SetTitle(text)
	dst_dropdown.SetTitle(text)
}

func SrcDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(dst_dropdown)
	case tcell.KeyEsc:
		pages.HidePage("lang_page")
	}
}

func DstDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTAB:
		app.SetFocus(src_dropdown)
	case tcell.KeyEsc:
		pages.HidePage("lang_page")
	}
}
