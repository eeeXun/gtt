package main

import (
	"fmt"
	"os"

	"github.com/eeeXun/gtt/internal/translate"
	config "github.com/spf13/viper"
)

// Search XDG_CONFIG_HOME or $HOME/.config
func configInit() {
	var (
		defaultConfigPath string
		defaultConfig     = map[string]interface{}{
			"hide_below":                             false,
			"transparent":                            false,
			"theme":                                  "Gruvbox",
			"source.border_color":                    "red",
			"destination.border_color":               "blue",
			"source.language.apertiumtranslate":      "English",
			"destination.language.apertiumtranslate": "English",
			"source.language.argostranslate":         "English",
			"destination.language.argostranslate":    "English",
			"source.language.googletranslate":        "English",
			"destination.language.googletranslate":   "English",
			"source.language.reversotranslate":       "English",
			"destination.language.reversotranslate":  "English",
			"translator":                             "ArgosTranslate",
		}
	)

	config.SetConfigName("gtt")
	config.SetConfigType("yaml")
	if len(os.Getenv("XDG_CONFIG_HOME")) > 0 {
		defaultConfigPath = os.Getenv("XDG_CONFIG_HOME") + "/gtt"
		config.AddConfigPath(defaultConfigPath)
	} else {
		defaultConfigPath = os.Getenv("HOME") + "/.config/gtt"
	}
	config.AddConfigPath("$HOME/.config/gtt")

	// Create config file if not exists
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
		if missing {
			config.WriteConfig()
		}
	}

	// setup
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
	// set argument language
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
