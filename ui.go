package main

import (
	"github.com/gdamore/tcell/v2"
)

func ui_init() {
	src_box.SetBorder(true).
		SetTitle(translator.src_lang).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.title_color).
		SetBackgroundColor(window.src.background_color)
	src_box.SetTextStyle(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color))
	src_box.SetSelectedStyle(tcell.StyleDefault.
		Background(window.src.selected_color).
		Foreground(window.src.foreground_color))
	src_box.SetInputCapture(InputHandle)

	dst_box.SetBorder(true).
		SetTitle(translator.dest_lang).
		SetBorderColor(window.dst.border_color).
		SetTitleColor(window.dst.title_color).
		SetBackgroundColor(window.dst.background_color)
	dst_box.SetTextColor(window.dst.foreground_color)

	src_dropdown.SetOptions(Lang, nil)
	dst_dropdown.SetOptions(Lang, nil)
	// src_dropdown.SetOptions([]string{"a", "b"}, nil)
	// dst_dropdown.SetOptions([]string{"a", "b"}, nil)
	src_dropdown.SetBorder(true).
		SetTitle("Select an option (hit Enter): ")
	dst_dropdown.SetBorder(true).
		SetTitle("Select an option (hit Enter): ")
}

func InputHandle(event *tcell.EventKey) *tcell.EventKey {
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
		translator.PlaySound(translator.dest_lang, dst_box.GetText(false))
	case tcell.KeyCtrlT:
		dst_box.SetText("TTT")
	case tcell.KeyCtrlS:
		dst_box.SetText("SSS")
	}

	return event
}
