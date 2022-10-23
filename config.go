package main

import (
	"os"
)

// Search XDG_CONFIG_HOME or $HOME/.config
func configInit() {
	var defaultConfigPath string

	config.SetConfigName("gtt")
	config.SetConfigType("yaml")
	if len(os.Getenv("$XDG_CONFIG_HOME")) > 0 {
		defaultConfigPath = os.Getenv("$XDG_CONFIG_HOME") + "/gtt"
		config.AddConfigPath(defaultConfigPath)
	} else {
		defaultConfigPath = os.Getenv("HOME") + "/.config/gtt"
	}
	config.AddConfigPath("$HOME/.config/gtt")

	// create config file if not exists
	if err := config.ReadInConfig(); err != nil {
		config.Set("transparent", false)
		config.Set("theme", "Gruvbox")
		config.Set("source.language", "English")
		config.Set("destination.language", "Chinese (Traditional)")
		if _, err = os.Stat(defaultConfigPath); os.IsNotExist(err) {
			os.Mkdir(defaultConfigPath, os.ModePerm)
		}
		config.SafeWriteConfig()
	}

	// setup
	theme = config.GetString("theme")
	transparent = config.GetBool("transparent")
	translator.srcLang = config.GetString("source.language")
	translator.dstLang = config.GetString("destination.language")
}

// Check if need to modify config file when quit program
func updateConfig() {
	changed := false

	if config.GetString("theme") != theme {
		changed = true
		config.Set("theme", theme)
	}
	if config.GetBool("transparent") != transparent {
		changed = true
		config.Set("transparent", transparent)
	}
	if config.GetString("source.language") != translator.srcLang {
		changed = true
		config.Set("source.language", translator.srcLang)
	}
	if config.GetString("destination.language") != translator.dstLang {
		changed = true
		config.Set("destination.language", translator.dstLang)
	}

	if changed {
		config.WriteConfig()
	}
}
