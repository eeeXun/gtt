package main

import (
	"gtt/internal/color"
	"os"

	"github.com/spf13/viper"
)

var (
	// settings
	config    = viper.New()
	style     = color.NewStyle()
	hideBelow bool
	// default config
	defaultConfig = map[string]interface{}{
		"transparent":                 false,
		"theme":                       "Gruvbox",
		"source.borderColor":          "red",
		"destination.borderColor":     "blue",
		"source.google.language":      "English",
		"destination.google.language": "Chinese (Traditional)",
		"hide_below":                  false,
		"translator":                  "google",
	}
)

// Search XDG_CONFIG_HOME or $HOME/.config
func configInit() {
	var defaultConfigPath string

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
	switch config.GetString("translator") {
	case "google":
		translator = googleTranslate
		if len(*srcLangArg) > 0 {
			translator.SetSrcLang(*srcLangArg)
		} else {
			translator.SetSrcLang(config.GetString("source.google.language"))
		}
		if len(*dstLangArg) > 0 {
			translator.SetDstLang(*dstLangArg)
		} else {
			translator.SetDstLang(config.GetString("destination.google.language"))
		}
	}
	hideBelow = config.GetBool("hide_below")
	style.Theme = config.GetString("theme")
	style.Transparent = config.GetBool("transparent")
	style.SetSrcBorderColor(config.GetString("source.borderColor")).
		SetDstBorderColor(config.GetString("destination.borderColor"))
}

// Check if need to modify config file when quit program
func updateConfig() {
	changed := false

	// Source language is not passed in argument
	if len(*srcLangArg) == 0 &&
		config.GetString("source.google.language") != googleTranslate.GetSrcLang() {
		changed = true
		config.Set("source.google.language", googleTranslate.GetSrcLang())
	}
	// Destination language is not passed in argument
	if len(*dstLangArg) == 0 &&
		config.GetString("destination.google.language") != googleTranslate.GetDstLang() {
		changed = true
		config.Set("destination.google.language", googleTranslate.GetDstLang())
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
