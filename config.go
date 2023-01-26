package main

import (
	"fmt"
	"gtt/internal/color"
	"os"

	"github.com/spf13/viper"
)

var (
	// settings
	config    = viper.New()
	style     = color.NewStyle()
	hideBelow bool
)

// Search XDG_CONFIG_HOME or $HOME/.config
func configInit() {
	var (
		defaultConfigPath string
		defaultConfig     = map[string]interface{}{
			"transparent":                          false,
			"theme":                                "Gruvbox",
			"source.borderColor":                   "red",
			"destination.borderColor":              "blue",
			"source.googletranslate.language":      "English",
			"destination.googletranslate.language": "Chinese (Traditional)",
			"source.libretranslate.language":       "English",
			"destination.libretranslate.language":  "Chinese",
			"hide_below":                           false,
			"translator":                           "LibreTranslate",
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
	for t_str, t := range translators {
		t.SetSrcLang(config.GetString(fmt.Sprintf("source.%s.language", t_str)))
		t.SetDstLang(config.GetString(fmt.Sprintf("destination.%s.language", t_str)))
	}
	translator = translators[config.GetString("translator")]
	hideBelow = config.GetBool("hide_below")
	style.Theme = config.GetString("theme")
	style.Transparent = config.GetBool("transparent")
	style.SetSrcBorderColor(config.GetString("source.borderColor")).
		SetDstBorderColor(config.GetString("destination.borderColor"))
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
			if config.GetString(fmt.Sprintf("source.%s.language", t_str)) != t.GetSrcLang() {
				changed = true
				config.Set(fmt.Sprintf("source.%s.language", t_str), t.GetSrcLang())
			}
		}
	}
	if len(*dstLangArg) == 0 {
		for t_str, t := range translators {
			if config.GetString(fmt.Sprintf("destination.%s.language", t_str)) != t.GetDstLang() {
				changed = true
				config.Set(fmt.Sprintf("destination.%s.language", t_str), t.GetDstLang())
			}
		}
	}
	if config.GetBool("hide_below") != hideBelow {
		changed = true
		config.Set("hide_below", hideBelow)
	}
	if config.GetString("theme") != style.Theme {
		changed = true
		config.Set("theme", style.Theme)
	}
	if config.GetBool("transparent") != style.Transparent {
		changed = true
		config.Set("transparent", style.Transparent)
	}
	if config.GetString("source.borderColor") != style.SrcBorderStr() {
		changed = true
		config.Set("source.borderColor", style.SrcBorderStr())
	}
	if config.GetString("destination.borderColor") != style.DstBorderStr() {
		changed = true
		config.Set("destination.borderColor", style.DstBorderStr())
	}

	if changed {
		config.WriteConfig()
	}
}
