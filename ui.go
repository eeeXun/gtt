package main

import (
	"fmt"
	"strconv"

	"github.com/eeeXun/gtt/internal/style"
	"github.com/eeeXun/gtt/internal/translate"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Item struct {
	item       tview.Primitive
	fixedSize  int
	proportion int
	focus      bool
}

const (
	popOutWindowHeight int    = 20
	langStrMaxLength   int    = 32
	keyMapText         string = `[#%[1]s]<C-c>[-]
	Exit program.
[#%[1]s]<Esc>[-]
	Toggle pop out window.
[#%[1]s]<%[2]s>[-]
	Translate from source to destination window.
[#%[1]s]<%[3]s>[-]
	Swap language.
[#%[1]s]<%[4]s>[-]
	Clear all text in source of translation window.
[#%[1]s]<%[5]s>[-]
	Copy selected text.
[#%[1]s]<%[6]s>[-]
	Copy all text in source of translation window.
[#%[1]s]<%[7]s>[-]
	Copy all text in destination of translation window.
[#%[1]s]<%[8]s>[-]
	Play text to speech on source of translation window.
[#%[1]s]<%[9]s>[-]
	Play text to speech on destination of translation window.
[#%[1]s]<%[10]s>[-]
	Stop playing text to speech.
[#%[1]s]<%[11]s>[-]
	Toggle transparent.
[#%[1]s]<%[12]s>[-]
	Toggle Definition/Example & Part of speech.
[#%[1]s]<Tab>, <S-Tab>[-]
	Cycle through the pop out widget.
[#%[1]s]<1>, <2>, <3>[-]
	Switch pop out window.`
)

func updateTranslateWindow() {
	if uiStyle.HideBelow {
		translateWindow.RemoveItem(translateBelowWidget)
	} else {
		translateWindow.AddItem(translateBelowWidget, 0, 1, false)
	}
}

func updateBackgroundColor() {
	// input/output
	srcInput.SetTextStyle(tcell.StyleDefault.
		Background(uiStyle.BackgroundColor()).
		Foreground(uiStyle.ForegroundColor())).
		SetBackgroundColor(uiStyle.BackgroundColor())
	dstOutput.SetBackgroundColor(uiStyle.BackgroundColor())
	defOutput.SetTextStyle(tcell.StyleDefault.
		Background(uiStyle.BackgroundColor()).
		Foreground(uiStyle.ForegroundColor())).
		SetBackgroundColor(uiStyle.BackgroundColor())
	posOutput.SetTextStyle(tcell.StyleDefault.
		Background(uiStyle.BackgroundColor()).
		Foreground(uiStyle.ForegroundColor())).
		SetBackgroundColor(uiStyle.BackgroundColor())

	// dropdown
	for _, dropdown := range []*tview.DropDown{
		translatorDropDown,
		srcLangDropDown,
		dstLangDropDown,
		themeDropDown,
		transparentDropDown,
		hideBelowDropDown,
		srcBorderDropDown,
		dstBorderDropDown} {
		dropdown.SetListStyles(tcell.StyleDefault.
			Background(uiStyle.BackgroundColor()).
			Foreground(uiStyle.ForegroundColor()),
			tcell.StyleDefault.
				Background(uiStyle.SelectedColor()).
				Foreground(uiStyle.PrefixColor())).
			SetBackgroundColor(uiStyle.BackgroundColor())
	}

	// key map
	keyMapMenu.SetBackgroundColor(uiStyle.BackgroundColor())
}

func updateBorderColor() {
	// input/output
	srcInput.SetBorderColor(uiStyle.SrcBorderColor()).
		SetTitleColor(uiStyle.SrcBorderColor())
	dstOutput.SetBorderColor(uiStyle.DstBorderColor()).
		SetTitleColor(uiStyle.DstBorderColor())
	defOutput.SetBorderColor(uiStyle.SrcBorderColor()).
		SetTitleColor(uiStyle.SrcBorderColor())
	posOutput.SetBorderColor(uiStyle.DstBorderColor()).
		SetTitleColor(uiStyle.DstBorderColor())

	// dropdown
	for _, srcDropDown := range []*tview.DropDown{srcLangDropDown, srcBorderDropDown} {
		srcDropDown.SetBorderColor(uiStyle.SrcBorderColor()).
			SetTitleColor(uiStyle.SrcBorderColor())
	}
	for _, dstDropDown := range []*tview.DropDown{dstLangDropDown, dstBorderDropDown} {
		dstDropDown.SetBorderColor(uiStyle.DstBorderColor()).
			SetTitleColor(uiStyle.DstBorderColor())
	}
}

func updateNonConfigColor() {
	// input/output
	srcInput.SetSelectedStyle(tcell.StyleDefault.
		Background(uiStyle.SelectedColor()).
		Foreground(uiStyle.ForegroundColor()))
	dstOutput.SetTextColor(uiStyle.ForegroundColor())
	defOutput.SetSelectedStyle(tcell.StyleDefault.
		Background(uiStyle.SelectedColor()).
		Foreground(uiStyle.ForegroundColor()))
	posOutput.SetSelectedStyle(tcell.StyleDefault.
		Background(uiStyle.SelectedColor()).
		Foreground(uiStyle.ForegroundColor()))

	// dropdown
	for _, noLabelDropDown := range []*tview.DropDown{srcLangDropDown, dstLangDropDown} {
		noLabelDropDown.SetFieldBackgroundColor(uiStyle.SelectedColor()).
			SetFieldTextColor(uiStyle.ForegroundColor()).
			SetPrefixTextColor(uiStyle.PrefixColor())
	}
	for _, labelDropDown := range []*tview.DropDown{
		translatorDropDown,
		themeDropDown,
		transparentDropDown,
		hideBelowDropDown,
		srcBorderDropDown,
		dstBorderDropDown} {
		labelDropDown.SetLabelColor(uiStyle.LabelColor()).
			SetFieldBackgroundColor(uiStyle.SelectedColor()).
			SetFieldTextColor(uiStyle.ForegroundColor()).
			SetPrefixTextColor(uiStyle.PrefixColor())
	}

	// button
	for _, button := range []*tview.Button{langButton, styleButton, keyMapButton} {
		button.SetStyle(tcell.StyleDefault.
			Background(uiStyle.SelectedColor()).
			Foreground(uiStyle.ForegroundColor())).
			SetActivatedStyle(
				tcell.StyleDefault.
					Background(uiStyle.PressColor()).
					Foreground(uiStyle.ForegroundColor()))
	}

	// key map
	keyMapMenu.SetTextColor(uiStyle.ForegroundColor()).
		SetText(fmt.Sprintf(keyMapText,
			fmt.Sprintf("%.6x", uiStyle.HighLightColor().TrueColor().Hex()),
			keyMaps["translate"],
			keyMaps["swap_language"],
			keyMaps["clear"],
			keyMaps["copy_selected"],
			keyMaps["copy_source"],
			keyMaps["copy_destination"],
			keyMaps["tts_source"],
			keyMaps["tts_destination"],
			keyMaps["stop_tts"],
			keyMaps["toggle_transparent"],
			keyMaps["toggle_below"],
		)).
		SetBorderColor(uiStyle.HighLightColor()).
		SetTitleColor(uiStyle.HighLightColor())
}

func updateAllColor() {
	updateBackgroundColor()
	updateBorderColor()
	updateNonConfigColor()
}

// SetSelectedFunc of DropDown need to update when options change
func updateLangDropDown() {
	srcLangDropDown.SetOptions(translator.GetAllLang(),
		func(text string, index int) {
			translator.SetSrcLang(text)
			srcInput.SetTitle(text)
			srcLangDropDown.SetTitle(text)
		})
	dstLangDropDown.SetOptions(translator.GetAllLang(),
		func(text string, index int) {
			translator.SetDstLang(text)
			dstOutput.SetTitle(text)
			dstLangDropDown.SetTitle(text)
		})
}

// Update language title and option
func updateCurrentLang() {
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

// If center is true, it will center the items
func attachItems(center bool, direction int, items ...Item) *tview.Flex {
	container := tview.NewFlex().SetDirection(direction)
	if center {
		container.AddItem(nil, 0, 1, false)
	}
	for _, item := range items {
		container.AddItem(item.item, item.fixedSize, item.proportion, item.focus)
	}
	if center {
		container.AddItem(nil, 0, 1, false)
	}
	return container
}

func showLangPopout() {
	mainPage.HidePage("stylePopOut")
	mainPage.HidePage("keyMapPopOut")
	mainPage.ShowPage("langPopOut")
	app.SetFocus(langCycle.GetCurrentUI())
}

func showStylePopout() {
	mainPage.HidePage("langPopOut")
	mainPage.HidePage("keyMapPopOut")
	mainPage.ShowPage("stylePopOut")
	app.SetFocus(styleCycle.GetCurrentUI())
}

func showKeyMapPopout() {
	mainPage.HidePage("langPopOut")
	mainPage.HidePage("stylePopOut")
	mainPage.ShowPage("keyMapPopOut")
}

func uiInit() {
	// pages
	mainPage.AddPage("translateWindow", translateWindow, true, true)
	mainPage.AddPage("langPopOut", langPopOut, true, false)
	mainPage.AddPage("stylePopOut", stylePopOut, true, false)
	mainPage.AddPage("keyMapPopOut", keyMapPopOut, true, false)

	// input/output
	srcInput.SetBorder(true)
	dstOutput.SetBorder(true)
	defOutput.SetBorder(true).SetTitle("Definition/Example")
	posOutput.SetBorder(true).SetTitle("Part of speech")

	// dropdown
	translatorDropDown.SetLabel("Translator: ").
		SetOptions(translate.AllTranslator, nil).
		SetCurrentOption(IndexOf(translator.GetEngineName(), translate.AllTranslator))
	srcLangDropDown.SetBorder(true)
	dstLangDropDown.SetBorder(true)
	themeDropDown.SetLabel("Theme: ").
		SetOptions(style.AllTheme, nil).
		SetCurrentOption(IndexOf(uiStyle.Theme, style.AllTheme))
	hideBelowDropDown.SetLabel("Hide below: ").
		SetOptions([]string{"true", "false"}, nil).
		SetCurrentOption(
			IndexOf(strconv.FormatBool(uiStyle.HideBelow),
				[]string{"true", "false"}))
	transparentDropDown.SetLabel("Transparent: ").
		SetOptions([]string{"true", "false"}, nil).
		SetCurrentOption(
			IndexOf(strconv.FormatBool(uiStyle.Transparent),
				[]string{"true", "false"}))
	srcBorderDropDown.SetLabel("Border Color: ").
		SetOptions(style.Palette, nil).
		SetCurrentOption(
			IndexOf(uiStyle.SrcBorderStr(),
				style.Palette)).
		SetBorder(true).
		SetTitle("Source")
	dstBorderDropDown.SetLabel("Border Color: ").
		SetOptions(style.Palette, nil).
		SetCurrentOption(
			IndexOf(uiStyle.DstBorderStr(),
				style.Palette)).
		SetBorder(true).
		SetTitle("Destination")

	// key map
	keyMapMenu.SetDynamicColors(true).
		SetBorder(true).
		SetTitle("Key Map")

	// window
	translateAboveWidget.SetDirection(tview.FlexColumn).
		AddItem(srcInput, 0, 1, true).
		AddItem(dstOutput, 0, 1, false)
	translateBelowWidget.SetDirection(tview.FlexColumn).
		AddItem(defOutput, 0, 1, false).
		AddItem(posOutput, 0, 1, false)
	translateWindow.SetDirection(tview.FlexRow).
		AddItem(translateAboveWidget, 0, 1, true)
	updateTranslateWindow()
	langPopOut.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(attachItems(true, tview.FlexColumn,
			Item{item: attachItems(false, tview.FlexRow,
				Item{item: attachItems(true, tview.FlexColumn,
					Item{item: attachItems(false, tview.FlexRow,
						Item{item: translatorDropDown, fixedSize: 0, proportion: 1, focus: false}),
						fixedSize: 0, proportion: 1, focus: false}),
					fixedSize: 1, proportion: 1, focus: false},
				Item{item: attachItems(false, tview.FlexColumn,
					Item{item: srcLangDropDown, fixedSize: 0, proportion: 1, focus: true},
					Item{item: dstLangDropDown, fixedSize: 0, proportion: 1, focus: false}),
					fixedSize: 0, proportion: 1, focus: true}),
				fixedSize: 2 * langStrMaxLength, proportion: 1, focus: true}),
			popOutWindowHeight, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)
	stylePopOut.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(attachItems(true, tview.FlexColumn,
			Item{item: attachItems(false, tview.FlexRow,
				Item{item: attachItems(true, tview.FlexColumn,
					Item{item: attachItems(false, tview.FlexRow,
						Item{item: themeDropDown, fixedSize: 0, proportion: 1, focus: true},
						Item{item: transparentDropDown, fixedSize: 0, proportion: 1, focus: false},
						Item{item: hideBelowDropDown, fixedSize: 0, proportion: 1, focus: false}),
						fixedSize: 0, proportion: 1, focus: true}),
					fixedSize: 3, proportion: 1, focus: true},
				Item{item: attachItems(false, tview.FlexColumn,
					Item{item: srcBorderDropDown, fixedSize: 0, proportion: 1, focus: false},
					Item{item: dstBorderDropDown, fixedSize: 0, proportion: 1, focus: false}),
					fixedSize: 0, proportion: 1, focus: false}),
				fixedSize: 2 * langStrMaxLength, proportion: 1, focus: true}),
			popOutWindowHeight, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)
	keyMapPopOut.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(attachItems(true, tview.FlexColumn,
			Item{item: keyMapMenu, fixedSize: 2 * langStrMaxLength, proportion: 1, focus: true}),
			popOutWindowHeight, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)

	updateAllColor()
	updateLangDropDown()
	updateCurrentLang()

	// handler
	app.SetInputCapture(appHandler)
	translateWindow.SetInputCapture(translateWindowHandler)
	for _, widget := range []*tview.TextArea{srcInput, defOutput, posOutput} {
		// fix for loop problem
		// https://github.com/golang/go/discussions/56010
		widget := widget
		widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			keyName := getKeyName(event)

			if len(keyName) == 0 {
				return event
			}

			switch keyName {
			case keyMaps["copy_selected"]:
				// copy selected text
				text, _, _ := widget.GetSelection()

				// only copy when text selected
				if len(text) > 0 {
					CopyToClipboard(text)
				}
				return nil
			}
			return event
		})
	}
	langPopOut.SetInputCapture(popOutHandler)
	stylePopOut.SetInputCapture(popOutHandler)
	keyMapPopOut.SetInputCapture(popOutHandler)
	translatorDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator = translators[text]
			updateLangDropDown()
			updateCurrentLang()
			SetTermTitle(translator.GetEngineName())
		})
	srcLangDropDown.SetDoneFunc(langDropDownHandler)
	dstLangDropDown.SetDoneFunc(langDropDownHandler)
	themeDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			uiStyle.Theme = text
			updateAllColor()
		})
	transparentDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			uiStyle.Transparent, _ = strconv.ParseBool(text)
			updateBackgroundColor()
		})
	hideBelowDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			uiStyle.HideBelow, _ = strconv.ParseBool(text)
			updateTranslateWindow()
		})
	srcBorderDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			uiStyle.SetSrcBorderColor(text)
			updateBorderColor()
		})
	dstBorderDropDown.SetDoneFunc(styleDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			uiStyle.SetDstBorderColor(text)
			updateBorderColor()
		})
	keyMapMenu.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEsc:
			mainPage.HidePage("keyMapPopOut")
		}
	})
	langButton.SetSelectedFunc(showLangPopout)
	styleButton.SetSelectedFunc(showStylePopout)
	keyMapButton.SetSelectedFunc(showKeyMapPopout)
}

func appHandler(event *tcell.EventKey) *tcell.EventKey {
	keyName := getKeyName(event)

	if len(keyName) == 0 {
		return event
	}

	switch keyName {
	case keyMaps["exit"]:
		app.Stop()
		return nil
	case keyMaps["toggle_transparent"]:
		// Toggle transparent
		uiStyle.Transparent = !uiStyle.Transparent
		// The following will trigger transparentDropDown SetDoneFunc
		transparentDropDown.SetCurrentOption(
			IndexOf(strconv.FormatBool(uiStyle.Transparent),
				[]string{"true", "false"}))
		return nil
	case keyMaps["toggle_below"]:
		// Toggle Hide below window
		uiStyle.HideBelow = !uiStyle.HideBelow
		// The following will trigger hideBelowDropDown SetDoneFunc
		hideBelowDropDown.SetCurrentOption(
			IndexOf(strconv.FormatBool(uiStyle.HideBelow),
				[]string{"true", "false"}))
		return nil
	}

	// Force C-c not to exit program
	if event.Key() == tcell.KeyCtrlC {
		return tcell.NewEventKey(tcell.KeyCtrlC, 0, tcell.ModNone)
	}

	return event
}

func translateWindowHandler(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEsc {
		showLangPopout()
		return nil
	}

	keyName := getKeyName(event)

	switch keyName {
	case keyMaps["translate"]:
		message := srcInput.GetText()
		// Only translate when message exist
		if len(message) > 0 {
			translation, err := translator.Translate(message)
			if err != nil {
				dstOutput.SetText(err.Error())
			} else {
				dstOutput.SetText(translation.TEXT)
				defOutput.SetText(translation.DEF, false)
				posOutput.SetText(translation.POS, false)
			}
		}
		return nil
	case keyMaps["clear"]:
		srcInput.SetText("", true)
		return nil
	case keyMaps["copy_source"]:
		// copy all text in Input
		text := srcInput.GetText()

		// only copy when text exist
		if len(text) > 0 {
			CopyToClipboard(text)
		}
		return nil
	case keyMaps["copy_destination"]:
		// copy all text in Output
		text := dstOutput.GetText(false)

		// only copy when text exist
		if len(text) > 0 {
			CopyToClipboard(text[:len(text)-1])
		}
		return nil
	case keyMaps["swap_language"]:
		translator.SwapLang()
		updateCurrentLang()
		srcText := srcInput.GetText()
		dstText := dstOutput.GetText(false)
		if len(dstText) > 0 {
			// GetText of Box contains "\n" if it has words
			srcInput.SetText(dstText[:len(dstText)-1], true)
		} else {
			srcInput.SetText(dstText, true)
		}
		dstOutput.SetText(srcText)
		return nil
	case keyMaps["tts_source"]:
		// Play text to speech on source of translation window.
		if translator.LockAvailable() {
			message := srcInput.GetText()
			// Only play when message exist
			if len(message) > 0 {
				translator.AcquireLock()
				go func() {
					err := translator.PlayTTS(translator.GetSrcLang(), message)
					if err != nil {
						srcInput.SetText(err.Error(), true)
						app.Draw()
					}
				}()
			}

		}
		return nil
	case keyMaps["tts_destination"]:
		// Play text to speech on destination of translation window.
		if translator.LockAvailable() {
			message := dstOutput.GetText(false)
			// Only play when message exist
			if len(message) > 0 {
				translator.AcquireLock()
				go func() {
					err := translator.PlayTTS(translator.GetDstLang(), message)
					if err != nil {
						dstOutput.SetText(err.Error())
						app.Draw()
					}
				}()
			}
		}
		return nil
	case keyMaps["stop_tts"]:
		// Stop play sound
		translator.StopTTS()
		return nil
	}

	return event
}

func popOutHandler(event *tcell.EventKey) *tcell.EventKey {
	ch := event.Rune()

	switch ch {
	case '1':
		showLangPopout()
	case '2':
		showStylePopout()
	case '3':
		showKeyMapPopout()
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
		mainPage.HidePage("langPopOut")
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
		mainPage.HidePage("stylePopOut")
	}
}
