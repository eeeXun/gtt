# GTT

Google Translate TUI

## ScreenShot

![screenshot](https://i.imgur.com/ECtL7ac.gif)

## Install

```
go get && go build
```

## Language in argument

You can pass `-src` and `-dst` in argument to set source and destination language.

```
gtt -src "English" -dst "Chinese (Traditional)"
```

See language on [Google Language support](https://cloud.google.com/translate/docs/languages)

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

`<C-y>`
Copy selected text in left window.

`<C-g>`
Copy all text in left window.

`<C-r>`
Copy all text in right window.

`<C-o>`
Play sound on left window.

`<C-p>`
Play sound on right window.

`<C-x>`
Stop play sound.

`<C-t>`
Toggle transparent.

`<C-\>`
Toggle Definition & Part of speech

`<Tab>`, `<S-Tab>`
Cycle through the pop out widget.

`<1>`, `<2>`, `<3>`
Switch pop out window.

## Dependencies

`xclip` For Linux to copy text.

`pbcopy` For macOS to copy text.

## Credit

[soimort/translate-shell](https://github.com/soimort/translate-shell) For translation URL.

[snsd0805/GoogleTranslate-TUI](https://github.com/snsd0805/GoogleTranslate-TUI) For inspiration.

[turk/free-google-translate](https://github.com/turk/free-google-translate) For Google translate in Golang.
