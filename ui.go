package main

import (
	"github.com/gdamore/tcell/v2"
)

func ui_init() {
	// page
	translate_page.SetInputCapture(TranslatePageHandler)

	// box
	src_box.SetBorder(true).
		SetTitle(translator.src_lang).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color).
		SetBackgroundColor(window.src.background_color)
	src_box.SetTextStyle(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color))
	src_box.SetSelectedStyle(tcell.StyleDefault.
		Background(window.src.selected_color).
		Foreground(window.src.foreground_color))
	src_box.SetInputCapture(SrcBoxHandler)

	dst_box.SetBorder(true).
		SetTitle(translator.dst_lang).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.border_color).
		SetBackgroundColor(window.dst.background_color)
	dst_box.SetTextColor(window.dst.foreground_color)

	// dropdown
	src_dropdown.SetOptions(Lang, nil).
		SetCurrentOption(IndexOf(translator.src_lang, Lang))
	src_dropdown.SetListStyles(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color),
		tcell.StyleDefault.
			Background(window.src.selected_color).
			Foreground(window.src.prefix_color))
	src_dropdown.SetFieldBackgroundColor(window.src.selected_color).
		SetFieldTextColor(window.src.foreground_color).
		SetPrefixTextColor(window.dst.prefix_color)
	src_dropdown.SetBorder(true).
		SetTitle(translator.src_lang).
		SetBackgroundColor(window.src.background_color).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color)
	src_dropdown.SetSelectedFunc(SrcSelected)
	src_dropdown.SetDoneFunc(SrcDropDownHandler)

	dst_dropdown.SetOptions(Lang, nil).
		SetCurrentOption(IndexOf(translator.dst_lang, Lang))
	dst_dropdown.SetListStyles(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color),
		tcell.StyleDefault.
			Background(window.src.selected_color).
			Foreground(window.src.prefix_color))
	dst_dropdown.SetFieldBackgroundColor(window.src.selected_color).
		SetFieldTextColor(window.src.foreground_color).
		SetPrefixTextColor(window.dst.prefix_color)
	dst_dropdown.SetBorder(true).
		SetTitle(translator.dst_lang).
		SetBackgroundColor(window.dst.background_color).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.border_color)
	dst_dropdown.SetSelectedFunc(DstSelected)
	dst_dropdown.SetDoneFunc(DstDropDownHandler)
}

func TranslatePageHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyEsc:
		pages.ShowPage("lang_page")
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

func SrcBoxHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	// panic(event.Name())

	switch key {
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
	case tcell.KeyCtrlT:
		dst_box.SetText("TTT")
	case tcell.KeyCtrlS:
		dst_box.SetText("SSS")
	}

	return event
}
