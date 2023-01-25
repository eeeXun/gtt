package main

import (
	"fmt"
	"gtt/internal/color"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	popOutWindowHeight int    = 20
	langStrMaxLength   int    = 32
	keyMapText         string = `[#%[1]s]<C-c>[-]
	Exit program.
[#%[1]s]<Esc>[-]
	Toggle pop out window.
[#%[1]s]<C-j>[-]
	Translate from source to destination window.
[#%[1]s]<C-s>[-]
	Swap language.
[#%[1]s]<C-q>[-]
	Clear all text in source of translation window.
[#%[1]s]<C-y>[-]
	Copy selected text.
[#%[1]s]<C-g>[-]
	Copy all text in source of translation window.
[#%[1]s]<C-r>[-]
	Copy all text in destination of translation window.
[#%[1]s]<C-o>[-]
	Play sound on source of translation window.
[#%[1]s]<C-p>[-]
	Play sound on destination of translation window.
[#%[1]s]<C-x>[-]
	Stop play sound.
[#%[1]s]<C-t>[-]
	Toggle transparent.
[#%[1]s]<C-\>[-]
	Toggle Definition & Part of speech
[#%[1]s]<Tab>, <S-Tab>[-]
	Cycle through the pop out widget.
[#%[1]s]<1>, <2>, <3>[-]
	Switch pop out window.`
)

func updateTranslateWindow() {
	translateWindow.Clear()
	if hideBelow {
		translateWindow.AddItem(translateAboveWidget, 0, 1, true)
	} else {
		translateWindow.SetDirection(tview.FlexRow).
			AddItem(translateAboveWidget, 0, 1, true).
			AddItem(translateBelowWidget, 0, 1, false)
	}
}

func updateBackgroundColor() {
	// input/output
	srcInput.SetTextStyle(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor())).
		SetBackgroundColor(style.BackgroundColor())
	dstOutput.SetBackgroundColor(style.BackgroundColor())
	defOutput.SetTextStyle(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor())).
		SetBackgroundColor(style.BackgroundColor())
	posOutput.SetTextStyle(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor())).
		SetBackgroundColor(style.BackgroundColor())

	// dropdown
	for _, dropdown := range []*tview.DropDown{
		srcLangDropDown,
		dstLangDropDown,
		themeDropDown,
		transparentDropDown,
		hideBelowDropDown,
		srcBorderDropDown,
		dstBorderDropDown} {
		dropdown.SetListStyles(tcell.StyleDefault.
			Background(style.BackgroundColor()).
			Foreground(style.ForegroundColor()),
			tcell.StyleDefault.
				Background(style.SelectedColor()).
				Foreground(style.PrefixColor())).
			SetBackgroundColor(style.BackgroundColor())
	}

	// key map
	keyMapMenu.SetBackgroundColor(style.BackgroundColor())
}

func updateBorderColor() {
	// input/output
	srcInput.SetBorderColor(style.SrcBorderColor()).
		SetTitleColor(style.SrcBorderColor())
	dstOutput.SetBorderColor(style.DstBorderColor()).
		SetTitleColor(style.DstBorderColor())
	defOutput.SetBorderColor(style.SrcBorderColor()).
		SetTitleColor(style.SrcBorderColor())
	posOutput.SetBorderColor(style.DstBorderColor()).
		SetTitleColor(style.DstBorderColor())

	// dropdown
	for _, srcDropDown := range []*tview.DropDown{srcLangDropDown, srcBorderDropDown} {
		srcDropDown.SetBorderColor(style.SrcBorderColor()).
			SetTitleColor(style.SrcBorderColor())
	}
	for _, dstDropDown := range []*tview.DropDown{dstLangDropDown, dstBorderDropDown} {
		dstDropDown.SetBorderColor(style.DstBorderColor()).
			SetTitleColor(style.DstBorderColor())
	}
}

func updateNonConfigColor() {
	// input/output
	srcInput.SetSelectedStyle(tcell.StyleDefault.
		Background(style.SelectedColor()).
		Foreground(style.ForegroundColor()))
	dstOutput.SetTextColor(style.ForegroundColor())
	defOutput.SetSelectedStyle(tcell.StyleDefault.
		Background(style.SelectedColor()).
		Foreground(style.ForegroundColor()))
	posOutput.SetSelectedStyle(tcell.StyleDefault.
		Background(style.SelectedColor()).
		Foreground(style.ForegroundColor()))

	// dropdown
	for _, noLabelDropDown := range []*tview.DropDown{srcLangDropDown, dstLangDropDown} {
		noLabelDropDown.SetFieldBackgroundColor(style.SelectedColor()).
			SetFieldTextColor(style.ForegroundColor()).
			SetPrefixTextColor(style.PrefixColor())
	}
	for _, labelDropDown := range []*tview.DropDown{
		themeDropDown,
		transparentDropDown,
		hideBelowDropDown,
		srcBorderDropDown,
		dstBorderDropDown} {
		labelDropDown.SetLabelColor(style.LabelColor()).
			SetFieldBackgroundColor(style.SelectedColor()).
			SetFieldTextColor(style.ForegroundColor()).
			SetPrefixTextColor(style.PrefixColor())
	}

	// button
	for _, button := range []*tview.Button{langButton, styleButton, keyMapButton} {
		button.SetLabelColor(style.ForegroundColor()).
			SetBackgroundColorActivated(style.PressColor()).
			SetLabelColorActivated(style.ForegroundColor()).
			SetBackgroundColor(style.SelectedColor())
	}

	// key map
	keyMapMenu.SetTextColor(style.ForegroundColor()).
		SetText(fmt.Sprintf(keyMapText,
			fmt.Sprintf("%.6x",
				style.HighLightColor().TrueColor().Hex()))).
		SetBorderColor(style.HighLightColor()).
		SetTitleColor(style.HighLightColor())
}

func updateAllColor() {
	updateBackgroundColor()
	updateBorderColor()
	updateNonConfigColor()
}

// Update title and option
func updateTitle() {
	srcInput.SetTitle(translator.GetSrcLang())
	dstOutput.SetTitle(translator.GetDstLang())
	srcLangDropDown.SetCurrentOption(
		IndexOf(translator.GetSrcLang(), translator.GetAllLang())).
		SetTitle(translator.GetSrcLang())
	dstLangDropDown.SetCurrentOption(
		IndexOf(translator.GetDstLang(), translator.GetAllLang())).
		SetTitle(translator.GetDstLang())
}

func attachButton() *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(langButton, 11, 1, true).
		AddItem(nil, 18, 1, false).
		AddItem(styleButton, 8, 1, true).
		AddItem(nil, 18, 1, false).
		AddItem(keyMapButton, 9, 1, true).
		AddItem(nil, 0, 1, false)
}

func uiInit() {
	// input/output
	srcInput.SetBorder(true)
	dstOutput.SetBorder(true)
	defOutput.SetBorder(true).SetTitle("Definition")
	posOutput.SetBorder(true).SetTitle("Part of speech")

	// dropdown
	for _, langDropDown := range []*tview.DropDown{srcLangDropDown, dstLangDropDown} {
		langDropDown.SetOptions(translator.GetAllLang(), nil).
			SetBorder(true)
	}
	themeDropDown.SetLabel("Theme: ").
		SetOptions(color.AllTheme, nil).
		SetCurrentOption(IndexOf(style.Theme, color.AllTheme))
	hideBelowDropDown.SetLabel("Hide below: ").
		SetOptions([]string{"true", "false"}, nil).
		SetCurrentOption(
			IndexOf(strconv.FormatBool(hideBelow),
				[]string{"true", "false"}))
	transparentDropDown.SetLabel("Transparent: ").
		SetOptions([]string{"true", "false"}, nil).
		SetCurrentOption(
			IndexOf(strconv.FormatBool(style.Transparent),
				[]string{"true", "false"}))
	srcBorderDropDown.SetLabel("Border Color: ").
		SetOptions(color.Palette, nil).
		SetCurrentOption(
			IndexOf(style.SrcBorderStr(),
				color.Palette)).
		SetBorder(true).
		SetTitle("Source")
	dstBorderDropDown.SetLabel("Border Color: ").
		SetOptions(color.Palette, nil).
		SetCurrentOption(
			IndexOf(style.DstBorderStr(),
				color.Palette)).
		SetBorder(true).
		SetTitle("Destination")

	// key map
	keyMapMenu.SetDynamicColors(true).
		SetText(fmt.Sprintf(keyMapText,
			fmt.Sprintf("%.6x", style.HighLightColor().TrueColor().Hex()))).
		SetBorder(true).
		SetTitle("Key Map")

	// window
	translateAboveWidget.SetDirection(tview.FlexColumn).
		AddItem(srcInput, 0, 1, true).
		AddItem(dstOutput, 0, 1, false)
	translateBelowWidget.SetDirection(tview.FlexColumn).
		AddItem(defOutput, 0, 1, false).
		AddItem(posOutput, 0, 1, false)
	updateTranslateWindow()
	langWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(srcLangDropDown, langStrMaxLength, 1, true).
			AddItem(dstLangDropDown, langStrMaxLength, 1, false).
			AddItem(nil, 0, 1, false),
			popOutWindowHeight, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)
	styleWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
					AddItem(nil, 0, 1, false).
					AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
						AddItem(themeDropDown, 0, 1, true).
						AddItem(transparentDropDown, 0, 1, false).
						AddItem(hideBelowDropDown, 0, 1, false),
						0, 1, true).
					AddItem(nil, 0, 1, false),
					3, 1, true).
				AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
					AddItem(srcBorderDropDown, 0, 1, false).
					AddItem(dstBorderDropDown, 0, 1, false),
					0, 1, false),
				2*langStrMaxLength, 1, true).
			AddItem(nil, 0, 1, false),
			popOutWindowHeight, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)
	keyMapWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(keyMapMenu, 2*langStrMaxLength, 1, true).
			AddItem(nil, 0, 1, false),
			popOutWindowHeight, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)

	updateAllColor()
	updateTitle()

	// handler
	mainPage.SetInputCapture(mainPageHandler)
	translateWindow.SetInputCapture(translateWindowHandler)
	for _, widget := range []*tview.TextArea{srcInput, defOutput, posOutput} {
		// fix for loop problem
		// https://github.com/golang/go/discussions/56010
		widget := widget
		widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			key := event.Key()
			switch key {
			case tcell.KeyCtrlY:
				// copy selected text
				text, _, _ := widget.GetSelection()

				// only copy when text selected
				if len(text) > 0 {
					CopyToClipboard(text)
				}
			}
			return event
		})
	}
	langWindow.SetInputCapture(popOutWindowHandler)
	styleWindow.SetInputCapture(popOutWindowHandler)
	keyMapWindow.SetInputCapture(popOutWindowHandler)
	srcLangDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator.SetSrcLang(text)
			srcInput.SetTitle(text)
			srcLangDropDown.SetTitle(text)
		})
	dstLangDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator.SetDstLang(text)
			dstOutput.SetTitle(text)
			dstLangDropDown.SetTitle(text)
		})
	themeDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			style.Theme = text
			updateAllColor()
		})
	transparentDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			style.Transparent, _ = strconv.ParseBool(text)
			updateBackgroundColor()
		})
	hideBelowDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			hideBelow, _ = strconv.ParseBool(text)
			updateTranslateWindow()
		})
	srcBorderDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			style.SetSrcBorderColor(text)
			updateBorderColor()
		})
	dstBorderDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			style.SetDstBorderColor(text)
			updateBorderColor()
		})
	keyMapMenu.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEsc:
			mainPage.HidePage("keyMapWindow")
		}
	})
	langButton.SetSelectedFunc(func() {
		mainPage.HidePage("styleWindow")
		mainPage.HidePage("keyMapWindow")
		mainPage.ShowPage("langWindow")
		app.SetFocus(langCycle.GetCurrentUI())
	})
	styleButton.SetSelectedFunc(func() {
		mainPage.HidePage("langWindow")
		mainPage.HidePage("keyMapWindow")
		mainPage.ShowPage("styleWindow")
		app.SetFocus(styleCycle.GetCurrentUI())
	})
	keyMapButton.SetSelectedFunc(func() {
		mainPage.HidePage("langWindow")
		mainPage.HidePage("styleWindow")
		mainPage.ShowPage("keyMapWindow")
	})
}

func mainPageHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyCtrlT:
		// Toggle transparent
		style.Transparent = !style.Transparent
		updateBackgroundColor()
		transparentDropDown.SetCurrentOption(
			IndexOf(strconv.FormatBool(style.Transparent),
				[]string{"true", "false"}))
	case tcell.KeyCtrlBackslash:
		hideBelow = !hideBelow
		updateTranslateWindow()
		hideBelowDropDown.SetCurrentOption(
			IndexOf(strconv.FormatBool(hideBelow),
				[]string{"true", "false"}))
	}

	return event
}

func translateWindowHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyEsc:
		mainPage.ShowPage("langWindow")
		app.SetFocus(langCycle.GetCurrentUI())
	case tcell.KeyCtrlJ:
		message := srcInput.GetText()
		// Only translate when message exist
		if len(message) > 0 {
			translation, definition, partOfSpeech, err := translator.Translate(message)
			if err != nil {
				dstOutput.SetText(err.Error())
			} else {
				dstOutput.SetText(translation)
				defOutput.SetText(definition, false)
				posOutput.SetText(partOfSpeech, false)
			}
		}
	case tcell.KeyCtrlQ:
		srcInput.SetText("", true)
	case tcell.KeyCtrlG:
		// copy all text in Input
		text := srcInput.GetText()

		// only copy when text exist
		if len(text) > 0 {
			CopyToClipboard(text)
		}
	case tcell.KeyCtrlR:
		// copy all text in Output
		text := dstOutput.GetText(false)

		// only copy when text exist
		if len(text) > 0 {
			CopyToClipboard(text[:len(text)-1])
		}
	case tcell.KeyCtrlS:
		translator.SwapLang()
		updateTitle()
		srcText := srcInput.GetText()
		dstText := dstOutput.GetText(false)
		if len(dstText) > 0 {
			// GetText of Box contains "\n" if it has words
			srcInput.SetText(dstText[:len(dstText)-1], true)
		} else {
			srcInput.SetText(dstText, true)
		}
		dstOutput.SetText(srcText)
	case tcell.KeyCtrlO:
		// Play source sound
		if translator.LockAvailable() {
			message := srcInput.GetText()
			// Only play when message exist
			if len(message) > 0 {
				translator.LockAcquire()
				go func() {
					err := translator.PlayTTS(translator.GetSrcLang(), message)
					if err != nil {
						srcInput.SetText(err.Error(), true)
					}
				}()
			}

		}
	case tcell.KeyCtrlP:
		// Play destination sound
		if translator.LockAvailable() {
			message := dstOutput.GetText(false)
			// Only play when message exist
			if len(message) > 0 {
				translator.LockAcquire()
				go func() {
					err := translator.PlayTTS(translator.GetDstLang(), message)
					if err != nil {
						dstOutput.SetText(err.Error())
					}
				}()
			}
		}
	case tcell.KeyCtrlX:
		// Stop play sound
		translator.StopTTS()
	}

	return event
}

func popOutWindowHandler(event *tcell.EventKey) *tcell.EventKey {
	ch := event.Rune()

	switch ch {
	case '1':
		mainPage.HidePage("styleWindow")
		mainPage.HidePage("keyMapWindow")
		mainPage.ShowPage("langWindow")
		app.SetFocus(langCycle.GetCurrentUI())
	case '2':
		mainPage.HidePage("langWindow")
		mainPage.HidePage("keyMapWindow")
		mainPage.ShowPage("styleWindow")
		app.SetFocus(styleCycle.GetCurrentUI())
	case '3':
		mainPage.HidePage("langWindow")
		mainPage.HidePage("styleWindow")
		mainPage.ShowPage("keyMapWindow")
	}

	return event
}

func langDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTab:
		langCycle.Increase()
		app.SetFocus(langCycle.GetCurrentUI())
	case tcell.KeyBacktab:
		langCycle.Decrease()
		app.SetFocus(langCycle.GetCurrentUI())
	case tcell.KeyEsc:
		mainPage.HidePage("langWindow")
	}
}

func styleDropDownHandler(key tcell.Key) {
	switch key {
	case tcell.KeyTab:
		styleCycle.Increase()
		app.SetFocus(styleCycle.GetCurrentUI())
	case tcell.KeyBacktab:
		styleCycle.Decrease()
		app.SetFocus(styleCycle.GetCurrentUI())
	case tcell.KeyEsc:
		mainPage.HidePage("styleWindow")
	}
}
