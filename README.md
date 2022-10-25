# GTT

Google Translate TUI

## ScreenShot

![screenshot](https://i.imgur.com/ECtL7ac.gif)

## Install

```
go get && go build
```

## Key Map

`<C-c>`
Exit program.

`<Esc>`
Toggle pop out window.

`<C-j>`
Translate from left window to right window.

`<C-s>`
Swap language.

`<C-q>`
Clear all text in left window.

`<C-l>`
Selected all text in left window.

`<A-s>`
Copy selected text in left window.

`<A-a>`
Copy all text in left window.

`<C-o>`
Play sound on left window.

`<C-p>`
Play sound on right window.

`<C-x>`
Stop play sound.

`<C-t>`
Toggle transparent.

`<Tab>`, `<S-Tab>`
Cycle through the pop out widget.

`<1>`, `<2>`, `<3>`
Switch pop out window.

## Dependencies

`xclip` For Linux to copy text.

`pbcopy` For macOS to copy text.

## Credit

[snsd0805/GoogleTranslate-TUI](https://github.com/snsd0805/GoogleTranslate-TUI) For inspiration.

[turk/free-google-translate](https://github.com/turk/free-google-translate) For Google translate in Golang.
