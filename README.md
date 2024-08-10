# gtt

Google Translate TUI (Originally)

Supported Translator:
[`Apertium`](https://www.apertium.org/),
[`Bing`](https://www.bing.com/translator),
[`ChatGPT`](https://chat.openai.com/),
[`DeepL`](https://deepl.com/translator)(only free API),
[`DeepLX`](https://github.com/OwO-Network/DeepLX),
[`Google`](https://translate.google.com/)(default),
[`Libre`](https://libretranslate.com/),
[`Reverso`](https://www.reverso.net/text-translation)

## ScreenShot

![screenshot](https://github.com/eeeXun/gtt/assets/58657914/3841c2bf-62f7-434a-9e77-91c3748c5675)

## ⚠️ Note for ChatGPT and DeepL

ChatGPT and DeepL translations require API keys, which can be obtained from
[OpenAI API keys](https://platform.openai.com/account/api-keys) and
[DeepL API signup](https://www.deepl.com/pro-api) pages, respectively. Note
that only the free API is supported for DeepL currently. Once you have your
API key add it to `$XDG_CONFIG_HOME/gtt/server.yaml` or `$HOME/.config/gtt/server.yaml`.
See the example in [server.yaml](example/server.yaml) file.

```yaml
api_key:
  chatgpt:
    value: CHATGPT_API_KEY # <- Replace with your API Key
    # file: $HOME/secrets/chatgpt.txt # <- You can also specify the file where to read API Key
  deepl:
    value: DEEPL_API_KEY # <- Replace with your API Key
    # file: $HOME/secrets/deepl.txt # <- You can also specify the file where to read API Key
```

## DeepLX

DeepLX is [self-hosted server](https://github.com/OwO-Network/DeepLX).
You must provide IP address and port at
`$XDG_CONFIG_HOME/gtt/server.yaml` or `$HOME/.config/gtt/server.yaml`.
The api key for DeepLX is optional, depending on your setting.
See the example in [server.yaml](example/server.yaml) file.

```yaml
api_key:
  deeplx:
    value: DEEPLX_API_KEY # <- Replace with your TOKEN
    # file: $HOME/secrets/deeplx.txt # <- You can also specify the file where to read API Key
host:
  deeplx: 127.0.0.1:1188 # <- Replace with your server IP address and port
```

## Libre

If you want to use official [LibreTranslate](https://libretranslate.com/), you have to obtain an API Key on their [website](https://portal.libretranslate.com/).
Alternatively, if you want to host it by yourself, you must provide the IP address and port.
Make sure add them to `$XDG_CONFIG_HOME/gtt/server.yaml` or `$HOME/.config/gtt/server.yaml`.
See the example in [server.yaml](example/server.yaml) file.

```yaml
api_key:
  libre:
    value: LIBRE_API_KEY # <- Replace with your API Key
    # file: $HOME/secrets/libre.txt # <- You can also specify the file where to read API Key
host:
  libre: 127.0.0.1:5000 # <- Replace with your server IP address and port
```

## Install

### Dependencies

For Arch Linux, you need `alsa-lib`.
For Ubuntu or Debian, you need `libasound2-dev`.
For RedHat-based Linux, you need `alsa-lib-devel`.

[`xclip`](https://github.com/astrand/xclip) (optional) - for Linux/X11 to copy text.

[`wl-clipboard`](https://github.com/bugaevc/wl-clipboard) (optional) - for Linux/Wayland to copy text.

Or, if your terminal supports OSC 52, you can enable OSC 52 in page 2 of the pop out menu to copy text.

### Arch Linux ([AUR](https://aur.archlinux.org/packages/gtt-bin))

```sh
yay -S gtt-bin
```

### Nix ❄️ ([nixpkgs-unstable](https://search.nixos.org/packages?channel=unstable&show=gtt&from=0&size=50&sort=relevance&type=packages&query=gtt))

add to your package list or run with:

```sh
nix-shell -p '(import <nixpkgs-unstable> {}).gtt' --run gtt
```

or with flakes enabled:

```sh
nix run github:nixos/nixpkgs#gtt
```

### Prebuild

Binary file is available in [Release Page](https://github.com/eeeXun/gtt/releases) for Linux and macOS on x86_64.

### From source

#### go install

```sh
go install -ldflags="-s -w" github.com/eeeXun/gtt@latest
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
Toggle pop out menu.

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
Switch pop out menu.

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

For key to combine with `Ctrl`, the value can be `C-Space`, `C-\`, `C-]`, `C-^`, `C-_` or `C-a` to `C-z`.

For key to combine with `Alt`, the value can be `A-Space` or `A-` with the character you want.

Or you can use function key, the value can be `F1` to `F64`.

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
- [Bing language-support](https://learn.microsoft.com/en-us/azure/cognitive-services/translator/language-support#translation) for `Bing`
- `ChatGPT` is same as `Google`. See [Google Language support](https://cloud.google.com/translate/docs/languages)
- [DeepL API docs](https://www.deepl.com/docs-api/translate-text/) for `DeepL`
- `DeepLX` is same as `DeepL`. See [DeepL API docs](https://cloud.google.com/translate/docs/languages)
- [Google Language support](https://cloud.google.com/translate/docs/languages) for `Google`
- [LibreTranslate Languages](https://libretranslate.com/languages) for `Libre`
- [Reverso Translation](https://www.reverso.net/text-translation) for `Reverso`

## Credit

[soimort/translate-shell](https://github.com/soimort/translate-shell),
[SimplyTranslate-Engines](https://codeberg.org/SimpleWeb/SimplyTranslate-Engines),
[s0ftik3/reverso-api](https://github.com/s0ftik3/reverso-api)
For request method.

[snsd0805/GoogleTranslate-TUI](https://github.com/snsd0805/GoogleTranslate-TUI) For inspiration.

[turk/free-google-translate](https://github.com/turk/free-google-translate) For Google translate in Golang.
