# gtt

Google Translate TUI (Originally)

Supported Translator:
[`Apertium`](https://www.apertium.org/),
[`Argos`](https://translate.argosopentech.com/),
[`Bing`](https://www.bing.com/translator),
[`ChatGPT`](https://chat.openai.com/),
[`Google`](https://translate.google.com/)(default),
[`Reverso`](https://www.reverso.net/text-translation)

## ⚠️ Note for ChatGPT

You need to apply an API key on [OpenAI API keys](https://platform.openai.com/account/api-keys).
And write it to `$XDG_CONFIG_HOME/gtt/gtt.yaml` or `$HOME/.config/gtt/gtt.yaml`.

```yaml
api_key:
    chatgpt: YOUR_API_KEY # <- Replace with your API Key
```

## ScreenShot

![screenshot](https://user-images.githubusercontent.com/58657914/213123592-5d8bccfb-ff80-4ad6-aaca-03b31c4c2c59.gif)

## Install

### Arch Linux ([AUR](https://aur.archlinux.org/packages/gtt-bin))

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

## Create a theme

You can create a theme with theme name. And you must provide the color of `bg`, `fg`, `gray`, `red`, `green`, `yellow`, `blue`, `purple`, `cyan`, `orange`.

And note that:

- `bg` is for background color
- `fg` is for foreground color
- `gray` is for selected color
- `yellow` is for label color
- `orange` is for KeyMap menu color

See the example in [theme.yaml](example/theme.yaml) file. This file should located under `$XDG_CONFIG_HOME/gtt/theme.yaml` or `$HOME/.config/gtt/theme.yaml`

## Language in argument

You can pass `-src` and `-dst` in argument to set source and destination language.

```
gtt -src "English" -dst "Chinese (Traditional)"
```

See available languages on:

- [Apertium Translate](https://www.apertium.org/) for `Apertium`
- [argosopentech/argos-translate](https://github.com/argosopentech/argos-translate#supported-languages) for `Argos`
- [Bing language-support](https://learn.microsoft.com/en-us/azure/cognitive-services/translator/language-support#translation) for `Bing`
- `ChatGPT` is same as `Google`. See [Google Language support](https://cloud.google.com/translate/docs/languages)
- [Google Language support](https://cloud.google.com/translate/docs/languages) for `Google`
- [Reverso Translation](https://www.reverso.net/text-translation) for `Reverso`

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
Play text to speech on source of translation window.

`<C-p>`
Play text to speech on destination of translation window.

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

## Credit

[soimort/translate-shell](https://github.com/soimort/translate-shell),
[SimplyTranslate-Engines](https://codeberg.org/SimpleWeb/SimplyTranslate-Engines),
[s0ftik3/reverso-api](https://github.com/s0ftik3/reverso-api)
For translation URL.

[snsd0805/GoogleTranslate-TUI](https://github.com/snsd0805/GoogleTranslate-TUI) For inspiration.

[turk/free-google-translate](https://github.com/turk/free-google-translate) For Google translate in Golang.
