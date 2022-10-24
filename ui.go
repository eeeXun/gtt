package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gtt/internal/color"
	"gtt/internal/translate"
	"strconv"
)

const (
	keyMapText string = `[#%[1]s]<C-c>[-]
	Exit program.
[#%[1]s]<Esc>[-]
	Toggle pop out window.
[#%[1]s]<C-j>[-]
	Translate from left window to right window.
[#%[1]s]<C-s>[-]
	Swap language.
[#%[1]s]<C-q>[-]
	Clear all text in left window.
[#%[1]s]<C-o>[-]
	Play sound on left window.
[#%[1]s]<C-p>[-]
	Play sound on right window.
[#%[1]s]<C-x>[-]
	Stop play sound.
[#%[1]s]<C-t>[-]
	Toggle transparent.
[#%[1]s]<Tab>, <S-Tab>[-]
	Cycle through the pop out widget.
[#%[1]s]<1>, <2>, <3>[-]
	Switch pop out window.`
)

func updateBackgroundColor() {
	// input/output
	srcInput.SetBackgroundColor(style.BackgroundColor())
	srcInput.SetTextStyle(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()))
	dstOutput.SetBackgroundColor(style.BackgroundColor())

	// dropdown
	srcLangDropDown.SetBackgroundColor(style.BackgroundColor())
	srcLangDropDown.SetListStyles(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()),
		tcell.StyleDefault.
			Background(style.SelectedColor()).
			Foreground(style.PrefixColor()))
	dstLangDropDown.SetBackgroundColor(style.BackgroundColor())
	dstLangDropDown.SetListStyles(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()),
		tcell.StyleDefault.
			Background(style.SelectedColor()).
			Foreground(style.PrefixColor()))
	themeDropDown.SetBackgroundColor(style.BackgroundColor())
	themeDropDown.SetListStyles(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()),
		tcell.StyleDefault.
			Background(style.SelectedColor()).
			Foreground(style.PrefixColor()))
	transparentDropDown.SetBackgroundColor(style.BackgroundColor())
	transparentDropDown.SetListStyles(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()),
		tcell.StyleDefault.
			Background(style.SelectedColor()).
			Foreground(style.PrefixColor()))
	srcBorderDropDown.SetBackgroundColor(style.BackgroundColor())
	srcBorderDropDown.SetListStyles(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()),
		tcell.StyleDefault.
			Background(style.SelectedColor()).
			Foreground(style.PrefixColor()))
	dstBorderDropDown.SetBackgroundColor(style.BackgroundColor())
	dstBorderDropDown.SetListStyles(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()),
		tcell.StyleDefault.
			Background(style.SelectedColor()).
			Foreground(style.PrefixColor()))

	// key map
	keyMapMenu.SetBackgroundColor(style.BackgroundColor())
}

func updateBorderColor() {
	// input/output
	srcInput.SetBorderColor(style.SrcBorderColor()).
		SetTitleColor(style.SrcBorderColor())
	dstOutput.SetBorderColor(style.DstBorderColor()).
		SetTitleColor(style.DstBorderColor())

	// dropdown
	srcLangDropDown.SetBorderColor(style.SrcBorderColor()).
		SetTitleColor(style.SrcBorderColor())
	dstLangDropDown.SetBorderColor(style.DstBorderColor()).
		SetTitleColor(style.DstBorderColor())
	srcBorderDropDown.SetBorderColor(style.SrcBorderColor()).
		SetTitleColor(style.SrcBorderColor())
	dstBorderDropDown.SetBorderColor(style.DstBorderColor()).
		SetTitleColor(style.DstBorderColor())
}

func updateNonConfigColor() {
	// input/output
	srcInput.SetSelectedStyle(tcell.StyleDefault.
		Background(style.SelectedColor()).
		Foreground(style.ForegroundColor()))
	dstOutput.SetTextColor(style.ForegroundColor())

	// dropdown
	srcLangDropDown.SetFieldBackgroundColor(style.SelectedColor()).
		SetFieldTextColor(style.ForegroundColor()).
		SetPrefixTextColor(style.PrefixColor())
	dstLangDropDown.SetFieldBackgroundColor(style.SelectedColor()).
		SetFieldTextColor(style.ForegroundColor()).
		SetPrefixTextColor(style.PrefixColor())
	themeDropDown.SetLabelColor(style.LabelColor())
	themeDropDown.SetFieldBackgroundColor(style.SelectedColor()).
		SetFieldTextColor(style.ForegroundColor()).
		SetPrefixTextColor(style.PrefixColor())
	transparentDropDown.SetLabelColor(style.LabelColor())
	transparentDropDown.SetFieldBackgroundColor(style.SelectedColor()).
		SetFieldTextColor(style.ForegroundColor()).
		SetPrefixTextColor(style.PrefixColor())
	srcBorderDropDown.SetLabelColor(style.LabelColor())
	srcBorderDropDown.SetFieldBackgroundColor(style.SelectedColor()).
		SetFieldTextColor(style.ForegroundColor()).
		SetPrefixTextColor(style.PrefixColor())
	dstBorderDropDown.SetLabelColor(style.LabelColor())
	dstBorderDropDown.SetFieldBackgroundColor(style.SelectedColor()).
		SetFieldTextColor(style.ForegroundColor()).
		SetPrefixTextColor(style.PrefixColor())

	// button
	langButton.SetLabelColor(style.ForegroundColor()).
		SetBackgroundColorActivated(style.PressColor()).
		SetLabelColorActivated(style.ForegroundColor()).
		SetBackgroundColor(style.SelectedColor())
	styleButton.SetLabelColor(style.ForegroundColor()).
		SetBackgroundColorActivated(style.PressColor()).
		SetLabelColorActivated(style.ForegroundColor()).
		SetBackgroundColor(style.SelectedColor())
	keyMapButton.SetLabelColor(style.ForegroundColor()).
		SetBackgroundColorActivated(style.PressColor()).
		SetLabelColorActivated(style.ForegroundColor()).
		SetBackgroundColor(style.SelectedColor())

	// key map
	keyMapMenu.SetTextColor(style.ForegroundColor()).
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
	srcInput.SetTitle(translator.SrcLang)
	dstOutput.SetTitle(translator.DstLang)
	srcLangDropDown.SetCurrentOption(IndexOf(translator.SrcLang, translate.Lang))
	srcLangDropDown.SetTitle(translator.SrcLang)
	dstLangDropDown.SetCurrentOption(IndexOf(translator.DstLang, translate.Lang))
	dstLangDropDown.SetTitle(translator.DstLang)
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

	// dropdown
	srcLangDropDown.SetBorder(true)
	srcLangDropDown.SetOptions(translate.Lang, nil)
	dstLangDropDown.SetBorder(true)
	dstLangDropDown.SetOptions(translate.Lang, nil)
	themeDropDown.SetLabel("Theme: ").
		SetOptions(color.AllTheme, nil).
		SetCurrentOption(IndexOf(style.Theme, color.AllTheme))
	transparentDropDown.SetLabel("Transparent: ").
		SetOptions([]string{"true", "false"}, nil).
		SetCurrentOption(
			IndexOf(strconv.FormatBool(style.Transparent),
				[]string{"true", "false"}))
	srcBorderDropDown.SetLabel("Border Color: ").
		SetOptions(color.Palette, nil).
		SetCurrentOption(IndexOf(style.SrcBorderStr(), color.Palette))
	srcBorderDropDown.SetBorder(true).
		SetTitle("Source")
	dstBorderDropDown.SetLabel("Border Color: ").
		SetOptions(color.Palette, nil).
		SetCurrentOption(IndexOf(style.DstBorderStr(), color.Palette))
	dstBorderDropDown.SetBorder(true).
		SetTitle("Destination")

	// key map
	keyMapMenu.SetBorder(true).
		SetTitle("Key Map")
	keyMapMenu.SetDynamicColors(true).
		SetText(fmt.Sprintf(keyMapText,
			fmt.Sprintf("%.6x", style.HighLightColor().TrueColor().Hex())))

	// window
	translateWindow.SetDirection(tview.FlexColumn).
		AddItem(srcInput, 0, 1, true).
		AddItem(dstOutput, 0, 1, false)
	langWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(srcLangDropDown, 32, 1, true).
			AddItem(dstLangDropDown, 32, 1, false).
			AddItem(nil, 0, 1, false),
			20, 1, true).
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
						AddItem(themeDropDown, 1, 1, true).
						AddItem(transparentDropDown, 1, 1, false),
						0, 1, true).
					AddItem(nil, 0, 1, false),
					2, 1, true).
				AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
					AddItem(srcBorderDropDown, 32, 1, false).
					AddItem(dstBorderDropDown, 32, 1, false),
					0, 1, false),
				64, 1, true).
			AddItem(nil, 0, 1, false),
			20, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)
	keyMapWindow.SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(keyMapMenu, 64, 1, true).
			AddItem(nil, 0, 1, false),
			20, 1, true).
		AddItem(attachButton(), 1, 1, false).
		AddItem(nil, 0, 1, false)

	updateAllColor()
	updateTitle()

	// handler
	mainPage.SetInputCapture(mainPageHandler)
	translateWindow.SetInputCapture(translatePageHandler)
	langWindow.SetInputCapture(popOutWindowHandler)
	styleWindow.SetInputCapture(popOutWindowHandler)
	keyMapWindow.SetInputCapture(popOutWindowHandler)
	srcLangDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator.SrcLang = text
			srcInput.SetTitle(text)
			srcLangDropDown.SetTitle(text)
		})
	dstLangDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator.DstLang = text
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
			mainPage.HidePage("keyMapPage")
		}
	})
	langButton.SetSelectedFunc(func() {
		mainPage.HidePage("stylePage")
		mainPage.HidePage("keyMapPage")
		mainPage.ShowPage("langPage")
		app.SetFocus(langCycle.GetCurrentUI())
	})
	styleButton.SetSelectedFunc(func() {
		mainPage.HidePage("langPage")
		mainPage.HidePage("keyMapPage")
		mainPage.ShowPage("stylePage")
		app.SetFocus(styleCycle.GetCurrentUI())
	})
	keyMapButton.SetSelectedFunc(func() {
		mainPage.HidePage("langPage")
		mainPage.HidePage("stylePage")
		mainPage.ShowPage("keyMapPage")
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
	}

	return event
}

func translatePageHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyEsc:
		mainPage.ShowPage("langPage")
		app.SetFocus(langCycle.GetCurrentUI())
	case tcell.KeyCtrlJ:
		message := srcInput.GetText()
		if len(message) > 0 {
			// Only translate when message exist
			result, err := translator.Translate(message)
			if err != nil {
				dstOutput.SetText(err.Error())
			} else {
				dstOutput.SetText(result)
			}
		}
	case tcell.KeyCtrlQ:
		srcInput.SetText("", true)
	case tcell.KeyCtrlS:
		translator.SrcLang, translator.DstLang = translator.DstLang, translator.SrcLang
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
		if translator.SoundLock.Available() {
			message := srcInput.GetText()
			if len(message) > 0 {
				// Only play when message exist
				translator.SoundLock.Acquire()
				go func() {
					err := translator.PlaySound(translator.SrcLang, message)
					if err != nil {
						srcInput.SetText(err.Error(), true)
					}
				}()
			}

		}
	case tcell.KeyCtrlP:
		// Play destination sound
		if translator.SoundLock.Available() {
			message := dstOutput.GetText(false)
			if len(message) > 0 {
				// Only play when message exist
				translator.SoundLock.Acquire()
				go func() {
					err := translator.PlaySound(translator.DstLang, message)
					if err != nil {
						dstOutput.SetText(err.Error())
					}
				}()
			}
		}
	case tcell.KeyCtrlX:
		// Stop play sound
		translator.SoundLock.Stop = true
	}

	return event
}

func popOutWindowHandler(event *tcell.EventKey) *tcell.EventKey {
	ch := event.Rune()

	switch ch {
	case '1':
		mainPage.HidePage("stylePage")
		mainPage.HidePage("keyMapPage")
		mainPage.ShowPage("langPage")
		app.SetFocus(langCycle.GetCurrentUI())
	case '2':
		mainPage.HidePage("langPage")
		mainPage.HidePage("keyMapPage")
		mainPage.ShowPage("stylePage")
		app.SetFocus(styleCycle.GetCurrentUI())
	case '3':
		mainPage.HidePage("langPage")
		mainPage.HidePage("stylePage")
		mainPage.ShowPage("keyMapPage")
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
		mainPage.HidePage("langPage")
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
		mainPage.HidePage("stylePage")
	}
}
