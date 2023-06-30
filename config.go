package main

import (
	"fmt"
	"os"

	"github.com/eeeXun/gtt/internal/style"
	"github.com/eeeXun/gtt/internal/translate"
	"github.com/spf13/viper"
)

var (
	// Main config
	config = viper.New()
)

// Search XDG_CONFIG_HOME or $HOME/.config
func configInit() {
	var (
		defaultConfigPath string
		themeConfig       = viper.New()
		keyMapConfig      = viper.New()
		defaultKeyMaps    = map[string]string{
			"exit":               "C-c",
			"translate":          "C-j",
			"swap_language":      "C-s",
			"clear":              "C-q",
			"copy_selected":      "C-y",
			"copy_source":        "C-g",
			"copy_destination":   "C-r",
			"tts_source":         "C-o",
			"tts_destination":    "C-p",
			"stop_tts":           "C-x",
			"toggle_transparent": "C-t",
			"toggle_below":       "C-\\",
		}
		defaultConfig = map[string]interface{}{
			"hide_below":                    false,
			"transparent":                   false,
			"theme":                         "gruvbox",
			"source.border_color":           "red",
			"destination.border_color":      "blue",
			"source.language.apertium":      "English",
			"destination.language.apertium": "English",
			"source.language.argos":         "English",
			"destination.language.argos":    "English",
			"source.language.bing":          "English",
			"destination.language.bing":     "English",
			"source.language.chatgpt":       "English",
			"destination.language.chatgpt":  "English",
			"source.language.deepl":         "English",
			"destination.language.deepl":    "English",
			"source.language.google":        "English",
			"destination.language.google":   "English",
			"source.language.reverso":       "English",
			"destination.language.reverso":  "English",
			"translator":                    "Google",
		}
	)

	config.SetConfigName("gtt")
	themeConfig.SetConfigName("theme")
	keyMapConfig.SetConfigName("keymap")
	for _, c := range []*viper.Viper{config, themeConfig, keyMapConfig} {
		c.SetConfigType("yaml")
	}
	if len(os.Getenv("XDG_CONFIG_HOME")) > 0 {
		defaultConfigPath = os.Getenv("XDG_CONFIG_HOME") + "/gtt"
		for _, c := range []*viper.Viper{config, themeConfig, keyMapConfig} {
			c.AddConfigPath(defaultConfigPath)
		}
	} else {
		defaultConfigPath = os.Getenv("HOME") + "/.config/gtt"
	}
	for _, c := range []*viper.Viper{config, themeConfig, keyMapConfig} {
		c.AddConfigPath("$HOME/.config/gtt")
	}

	// Import theme if file exists
	if err := themeConfig.ReadInConfig(); err == nil {
		var (
			palate = make(map[string]int32)
			colors = []string{"bg", "fg", "gray", "red", "green", "yellow", "blue", "purple", "cyan", "orange"}
		)
		for name := range themeConfig.AllSettings() {
			for _, color := range colors {
				palate[color] = themeConfig.GetInt32(fmt.Sprintf("%s.%s", name, color))
			}
			style.NewTheme(name, palate)
		}
	}
	// Create config file if it does not exist
	// Otherwise check if config value is missing
	if err := config.ReadInConfig(); err != nil {
		for key, value := range defaultConfig {
			config.Set(key, value)
		}
		if _, err = os.Stat(defaultConfigPath); os.IsNotExist(err) {
			os.MkdirAll(defaultConfigPath, os.ModePerm)
		}
		config.SafeWriteConfig()
	} else {
		missing := false
		for key, value := range defaultConfig {
			if config.Get(key) == nil {
				config.Set(key, value)
				missing = true
			}
		}
		// Set to default theme if theme in config does not exist
		if IndexOf(config.GetString("theme"), style.AllTheme) < 0 {
			config.Set("theme", defaultConfig["theme"])
			missing = true
		}
		// Set to default translator if translator in config does not exist
		if IndexOf(config.GetString("translator"), translate.AllTranslator) < 0 {
			config.Set("translator", defaultConfig["translator"])
			missing = true
		}
		if missing {
			config.WriteConfig()
		}
	}

	// Setup key map
	// If keymap file exist and action in file exist, then set the keyMap
	// Otherwise, set to defaultKeyMap
	if err := keyMapConfig.ReadInConfig(); err == nil {
		for action, key := range defaultKeyMaps {
			if keyMapConfig.Get(action) == nil {
				keyMaps[action] = key
			} else {
				keyMaps[action] = keyMapConfig.GetString(action)
			}
		}
	} else {
		for action, key := range defaultKeyMaps {
			keyMaps[action] = key
		}
	}
	// Setup
	for _, name := range translate.AllTranslator {
		translators[name] = translate.NewTranslator(name)
		translators[name].SetSrcLang(
			config.GetString(fmt.Sprintf("source.language.%s", name)))
		translators[name].SetDstLang(
			config.GetString(fmt.Sprintf("destination.language.%s", name)))
	}
	translator = translators[config.GetString("translator")]
	uiStyle.Theme = config.GetString("theme")
	uiStyle.HideBelow = config.GetBool("hide_below")
	uiStyle.Transparent = config.GetBool("transparent")
	uiStyle.SetSrcBorderColor(config.GetString("source.border_color")).
		SetDstBorderColor(config.GetString("destination.border_color"))
	// Set API Keys
	for _, name := range []string{"ChatGPT", "DeepL"} {
		if config.Get(fmt.Sprintf("api_key.%s", name)) != nil {
			translators[name].SetAPIKey(config.GetString(fmt.Sprintf("api_key.%s", name)))
		}
	}
	// Set argument language
	if len(*srcLangArg) > 0 {
		translator.SetSrcLang(*srcLangArg)
	}
	if len(*dstLangArg) > 0 {
		translator.SetDstLang(*dstLangArg)
	}
}

// Check if need to modify config file when quit program
func updateConfig() {
	changed := false

	// Source language is not passed in argument
	if len(*srcLangArg) == 0 {
		for t_str, t := range translators {
			if config.GetString(fmt.Sprintf("source.language.%s", t_str)) != t.GetSrcLang() {
				changed = true
				config.Set(fmt.Sprintf("source.language.%s", t_str), t.GetSrcLang())
			}
		}
	}
	// Destination language is not passed in argument
	if len(*dstLangArg) == 0 {
		for t_str, t := range translators {
			if config.GetString(fmt.Sprintf("destination.language.%s", t_str)) != t.GetDstLang() {
				changed = true
				config.Set(fmt.Sprintf("destination.language.%s", t_str), t.GetDstLang())
			}
		}
	}
	if config.GetString("translator") != translator.GetEngineName() {
		changed = true
		config.Set("translator", translator.GetEngineName())
	}
	if config.GetBool("hide_below") != uiStyle.HideBelow {
		changed = true
		config.Set("hide_below", uiStyle.HideBelow)
	}
	if config.GetString("theme") != uiStyle.Theme {
		changed = true
		config.Set("theme", uiStyle.Theme)
	}
	if config.GetBool("transparent") != uiStyle.Transparent {
		changed = true
		config.Set("transparent", uiStyle.Transparent)
	}
	if config.GetString("source.border_color") != uiStyle.SrcBorderStr() {
		changed = true
		config.Set("source.border_color", uiStyle.SrcBorderStr())
	}
	if config.GetString("destination.border_color") != uiStyle.DstBorderStr() {
		changed = true
		config.Set("destination.border_color", uiStyle.DstBorderStr())
	}

	if changed {
		config.WriteConfig()
	}
}
