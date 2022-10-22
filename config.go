package main

import (
	"os"
)

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
		config.Set("source_language", "English")
		config.Set("destination_language", "Chinese (Traditional)")
		if _, err = os.Stat(defaultConfigPath); os.IsNotExist(err) {
			os.Mkdir(defaultConfigPath, os.ModePerm)
		}
		config.SafeWriteConfig()
	}

	// setup
	theme = config.GetString("theme")
	transparent = config.GetBool("transparent")
	translator.src_lang = config.GetString("source_language")
	translator.dst_lang = config.GetString("destination_language")
}

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
	if config.GetString("source_language") != translator.src_lang {
		changed = true
		config.Set("source_language", translator.src_lang)
	}
	if config.GetString("destination_language") != translator.dst_lang {
		changed = true
		config.Set("destination_language", translator.dst_lang)
	}

	if changed {
		config.WriteConfig()
	}
}
