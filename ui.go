package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

type UICycle struct {
	widget []tview.Primitive
	index  int
	len    int
}

func NewUICycle(widgets ...tview.Primitive) *UICycle {
	var w []tview.Primitive

	for _, widget := range widgets {
		w = append(w, widget)
	}

	return &UICycle{
		widget: w,
		index:  0,
		len:    len(w),
	}
}

func (ui *UICycle) Increase() {
	ui.index = (ui.index + 1) % ui.len
}

func (ui *UICycle) Decrease() {
	ui.index = ((ui.index-1)%ui.len + ui.len) % ui.len
}

func (ui *UICycle) GetCurrentUI() tview.Primitive {
	return ui.widget[ui.index]
}

func updateBackgroundColor() {
	// box
	srcBox.SetBackgroundColor(style.BackgroundColor())
	srcBox.SetTextStyle(tcell.StyleDefault.
		Background(style.BackgroundColor()).
		Foreground(style.ForegroundColor()))
	dstBox.SetBackgroundColor(style.BackgroundColor())

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
}

func updateBorderColor() {
	// box
	srcBox.SetBorderColor(style.SrcBorderColor()).
		SetTitleColor(style.SrcBorderColor())
	dstBox.SetBorderColor(style.DstBorderColor()).
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
	// box
	srcBox.SetSelectedStyle(tcell.StyleDefault.
		Background(style.SelectedColor()).
		Foreground(style.ForegroundColor()))
	dstBox.SetTextColor(style.ForegroundColor())

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
	menuButton.SetLabelColor(style.ForegroundColor()).
		SetBackgroundColorActivated(style.PressColor()).
		SetLabelColorActivated(style.ForegroundColor()).
		SetBackgroundColor(style.SelectedColor())
}

func updateAllColor() {
	updateBackgroundColor()
	updateBorderColor()
	updateNonConfigColor()
}

// Update title and option
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
	srcBox.SetBorder(true)
	dstBox.SetBorder(true)

	// dropdown
	srcLangDropDown.SetBorder(true)
	srcLangDropDown.SetOptions(Lang, nil)
	dstLangDropDown.SetBorder(true)
	dstLangDropDown.SetOptions(Lang, nil)
	themeDropDown.SetLabel("Theme: ").
		SetOptions(AllTheme, nil).
		SetCurrentOption(IndexOf(style.Theme, AllTheme))
	transparentDropDown.SetLabel("Transparent: ").
		SetOptions([]string{"true", "false"}, nil).
		SetCurrentOption(
			IndexOf(strconv.FormatBool(style.Transparent),
				[]string{"true", "false"}))
	srcBorderDropDown.SetLabel("Border Color: ").
		SetOptions(Palette, nil).
		SetCurrentOption(IndexOf(style.SrcBorderStr(), Palette))
	srcBorderDropDown.SetBorder(true).
		SetTitle("Source")
	dstBorderDropDown.SetLabel("Border Color: ").
		SetOptions(Palette, nil).
		SetCurrentOption(IndexOf(style.DstBorderStr(), Palette))
	dstBorderDropDown.SetBorder(true).
		SetTitle("Destination")

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

	updateAllColor()
	updateTitle()

	// handler
	mainPage.SetInputCapture(pagesHandler)
	langWindow.SetInputCapture(langWindowHandler)
	styleWindow.SetInputCapture(styleWindowHandler)
	translateWindow.SetInputCapture(translatePageHandler)
	srcLangDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator.srcLang = text
			srcBox.SetTitle(text)
			srcLangDropDown.SetTitle(text)
		})
	dstLangDropDown.SetDoneFunc(langDropDownHandler).
		SetSelectedFunc(func(text string, index int) {
			translator.dstLang = text
			dstBox.SetTitle(text)
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
	langButton.SetSelectedFunc(func() {
		mainPage.HidePage("stylePage")
		mainPage.ShowPage("langPage")
		app.SetFocus(langCycle.GetCurrentUI())
	})
	styleButton.SetSelectedFunc(func() {
		mainPage.HidePage("langPage")
		mainPage.ShowPage("stylePage")
		app.SetFocus(styleCycle.GetCurrentUI())
	})
}

func pagesHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyCtrlT:
		style.Transparent = !style.Transparent
		updateBackgroundColor()
		transparentDropDown.SetCurrentOption(
			IndexOf(strconv.FormatBool(style.Transparent),
				[]string{"true", "false"}))
	}

	return event
}

func langWindowHandler(event *tcell.EventKey) *tcell.EventKey {
	ch := event.Rune()

	switch ch {
	case '2':
		mainPage.HidePage("langPage")
		mainPage.ShowPage("stylePage")
		app.SetFocus(styleCycle.GetCurrentUI())
	}

	return event
}

func styleWindowHandler(event *tcell.EventKey) *tcell.EventKey {
	ch := event.Rune()

	switch ch {
	case '1':
		mainPage.HidePage("stylePage")
		mainPage.ShowPage("langPage")
		app.SetFocus(langCycle.GetCurrentUI())
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
		// Play source sound
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
		// Play destination sound
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
		// Stop play sound
		translator.soundLock.stop = true
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
