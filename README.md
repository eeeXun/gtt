# gtt

Google Translate TUI (Originally)

Supported Translator:
[`Apertium`](https://www.apertium.org/),
[`Argos`](https://translate.argosopentech.com/),
[`Bing`](https://www.bing.com/translator),
[`ChatGPT`](https://chat.openai.com/),
[`DeepL`](https://deepl.com/translator)(only free API),
[`Google`](https://translate.google.com/)(default),
[`Reverso`](https://www.reverso.net/text-translation)

## ⚠️ Note for ChatGPT and DeepL

ChatGPT and DeepL translations require API keys, which can be obtained from
[OpenAI API keys](https://platform.openai.com/account/api-keys) and
[DeepL API signup](https://www.deepl.com/pro-api) pages, respectively. Note
that only the free API is supported for DeepL currently. Once you have your
API key add it to `$XDG_CONFIG_HOME/gtt/gtt.yaml` or
`$HOME/.config/gtt/gtt.yaml`

```yaml
api_key:
    chatgpt: CHATGPT_API_KEY # <- Replace with your API Key
    deepl: DEEPL_API_KEY # <- Replace with your API Key
```

## ScreenShot

![screenshot](https://github.com/eeeXun/gtt/assets/58657914/3841c2bf-62f7-434a-9e77-91c3748c5675)

## Install

### Dependencies

For Arch Linux, you need `alsa-lib`.
For Ubuntu or Debian, you need `libasound2-dev`.
For RedHat-based Linux, you need `alsa-lib-devel`.

[`xclip`](https://github.com/astrand/xclip) (optional) - for Linux/X11 to copy text.

[`wl-clipboard`](https://github.com/bugaevc/wl-clipboard) (optional) - for Linux/Wayland to copy text.

### Arch Linux ([AUR](https://aur.archlinux.org/packages/gtt-bin))

```sh
yay -S gtt-bin
```

### Prebuild

Binary file is available in [Release Page](https://github.com/eeeXun/gtt/releases) for Linux and macOS on x86_64.

### From source

#### go install

```sh
go install github.com/eeeXun/gtt@latest
```

And make sure `$HOME/go/bin` is in your `$PATH`

```sh
export PATH=$PATH:$HOME/go/bin
```

#### go build

```sh
git clone https://github.com/eeeXun/gtt.git && cd gtt && go build -ldflags="-s -w -X main.version=$(git describe --tags)"
```

### Run on Docker ([Docker Hub](https://hub.docker.com/r/eeexun/gtt/tags))

```sh
docker run -it eeexun/gtt:latest
```

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
Stop playing text to speech.

`<C-t>`
Toggle transparent.

`<C-\>`
Toggle Definition/Example & Part of speech.

`<Tab>`, `<S-Tab>`
Cycle through the pop out widget.

`<1>`, `<2>`, `<3>`
Switch pop out window.

### Customize key map

You can overwrite the following key

- `exit`: Exit program.
- `translate`: Translate from source to destination window.
- `swap_language`: Swap language.
- `clear`: Clear all text in source of translation window.
- `copy_selected`: Copy selected text.
- `copy_source`: Copy all text in source of translation window.
- `copy_destination`: Copy all text in destination of translation window.
- `tts_source`: Play text to speech on source of translation window.
- `tts_destination`: Play text to speech on destination of translation window.
- `stop_tts`: Stop playing text to speech.
- `toggle_transparent`: Toggle transparent.
- `toggle_below`: Toggle Definition/Example & Part of speech.

For key to combine with `Ctrl`, the value can be `"C-Space"`, `"C-\\"`, `"C-]"`, `"C-^"`, `"C-_"` or `"C-a"` to `"C-z"`.

For key to combine with `Alt`, the value can be `"A-Space"` or `"A-"` + the character you want.

Or you can use function key, the value can be `"F1"` to `"F64"`.

See the example in [keymap.yaml](example/keymap.yaml) file. This file should be located at `$XDG_CONFIG_HOME/gtt/keymap.yaml` or `$HOME/.config/gtt/keymap.yaml`.

## Customize theme

You can create a theme with theme name. And you must provide the color of `bg`, `fg`, `gray`, `red`, `green`, `yellow`, `blue`, `purple`, `cyan`, `orange`.

And note that:

- `bg` is for background color
- `fg` is for foreground color
- `gray` is for selected color
- `yellow` is for label color
- `orange` is for KeyMap menu color
- `purple` is for button pressed color

See the example in [theme.yaml](example/theme.yaml) file. This file should be located at `$XDG_CONFIG_HOME/gtt/theme.yaml` or `$HOME/.config/gtt/theme.yaml`.

## Language in argument

You can pass `-src` and `-dst` in argument to set source and destination language.

```sh
gtt -src "English" -dst "Chinese (Traditional)"
```

See available languages on:

- [Apertium Translate](https://www.apertium.org/) for `Apertium`
- [argosopentech/argos-translate](https://github.com/argosopentech/argos-translate#supported-languages) for `Argos`
- [Bing language-support](https://learn.microsoft.com/en-us/azure/cognitive-services/translator/language-support#translation) for `Bing`
- `ChatGPT` is same as `Google`. See [Google Language support](https://cloud.google.com/translate/docs/languages)
- [DeepL API docs](https://www.deepl.com/docs-api/translate-text/) for `DeepL`
- [Google Language support](https://cloud.google.com/translate/docs/languages) for `Google`
- [Reverso Translation](https://www.reverso.net/text-translation) for `Reverso`

## Credit

[soimort/translate-shell](https://github.com/soimort/translate-shell),
[SimplyTranslate-Engines](https://codeberg.org/SimpleWeb/SimplyTranslate-Engines),
[s0ftik3/reverso-api](https://github.com/s0ftik3/reverso-api)
For translation URL.

[snsd0805/GoogleTranslate-TUI](https://github.com/snsd0805/GoogleTranslate-TUI) For inspiration.

[turk/free-google-translate](https://github.com/turk/free-google-translate) For Google translate in Golang.
