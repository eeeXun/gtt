# gtt

Google Translate TUI (Originally)

Supported Translator:
[`ApertiumTranslate`](https://www.apertium.org/),
[`ArgosTranslate`](https://translate.argosopentech.com/),
[`GoogleTranslate`](https://translate.google.com/),
[`ReversoTranslate`](https://www.reverso.net/text-translation)

## ScreenShot

![screenshot](https://user-images.githubusercontent.com/58657914/213123592-5d8bccfb-ff80-4ad6-aaca-03b31c4c2c59.gif)

## Install

### Arch Linux (AUR)

```
yay -S gtt-bin
```

### Prebuild

Binary file is available in [Release Page](https://github.com/eeeXun/gtt/releases) for Linux and macOS on x86_64.

### From source

```
go install github.com/eeeXun/gtt@latest
```

### Run on Docker

```
docker run -it eeexun/gtt
```

## Language in argument

You can pass `-src` and `-dst` in argument to set source and destination language.

```
gtt -src "English" -dst "Chinese (Traditional)"
```

See language on:

- [Apertium Translate](https://www.apertium.org/) for `ApertiumTranslate`
- [argosopentech/argos-translate](https://github.com/argosopentech/argos-translate#supported-languages) for `ArgosTranslate`
- [Google Language support](https://cloud.google.com/translate/docs/languages) for `GoogleTranslate`
- [Reverso Translation](https://www.reverso.net/text-translation) for `ReversoTranslate`

## Key Map

`<C-c>`
Exit program.

`<Esc>`
Toggle pop out window.

`<C-j>`
Translate from source to destination window.

`<C-s>`
Swap language.

`<C-q>`
Clear all text in source of translation window.

`<C-y>`
Copy selected text.

`<C-g>`
Copy all text in source of translation window.

`<C-r>`
Copy all text in destination of translation window.

`<C-o>`
Play sound on source of translation window.

`<C-p>`
Play sound on destination of translation window.

`<C-x>`
Stop play sound.

`<C-t>`
Toggle transparent.

`<C-\>`
Toggle Definition/Example & Part of speech.

`<Tab>`, `<S-Tab>`
Cycle through the pop out widget.

`<1>`, `<2>`, `<3>`
Switch pop out window.

## Dependencies

[`xclip`](https://github.com/astrand/xclip) for Linux/X11 to copy text.

[`wl-clipboard`](https://github.com/bugaevc/wl-clipboard) for Linux/Wayland to copy text.

`pbcopy` For macOS to copy text.

## Credit

[soimort/translate-shell](https://github.com/soimort/translate-shell),
[SimplyTranslate-Engines](https://codeberg.org/SimpleWeb/SimplyTranslate-Engines),
[s0ftik3/reverso-api](https://github.com/s0ftik3/reverso-api)
For translation URL.

[snsd0805/GoogleTranslate-TUI](https://github.com/snsd0805/GoogleTranslate-TUI) For inspiration.

[turk/free-google-translate](https://github.com/turk/free-google-translate) For Google translate in Golang.
