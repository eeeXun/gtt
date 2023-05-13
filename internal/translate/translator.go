package translate

import (
	"github.com/eeeXun/gtt/internal/translate/apertium"
	"github.com/eeeXun/gtt/internal/translate/argos"
	"github.com/eeeXun/gtt/internal/translate/bing"
	"github.com/eeeXun/gtt/internal/translate/chatgpt"
	"github.com/eeeXun/gtt/internal/translate/core"
	"github.com/eeeXun/gtt/internal/translate/deepl"
	"github.com/eeeXun/gtt/internal/translate/google"
	"github.com/eeeXun/gtt/internal/translate/reverso"
)

var (
	AllTranslator = []string{
		"Apertium",
		"Argos",
		"Bing",
		"ChatGPT",
		"DeepL",
		"Google",
		"Reverso",
	}
)

type Translator interface {
	// Get engine name of the translator
	GetEngineName() string

	// Get all languages of the translator
	GetAllLang() []string

	// Get source language of the translator
	GetSrcLang() string

	// Get destination language of the translator
	GetDstLang() string

	// Set source language of the translator
	SetSrcLang(lang string)

	// Set destination language of the translator
	SetDstLang(lang string)

	// Swap source and destination language of the translator
	SwapLang()

	// Set API Key
	SetAPIKey(key string)

	// Check if lock is available
	LockAvailable() bool

	// Acquire the lock
	AcquireLock()

	// Stop text to speech
	StopTTS()

	// Translate from source to destination language
	Translate(message string) (translation *core.Translation, err error)

	// Play text to speech
	PlayTTS(lang, message string) error
}

func NewTranslator(name string) Translator {
	var translator Translator

	switch name {
	case "Apertium":
		translator = apertium.NewTranslator()
	case "Argos":
		translator = argos.NewTranslator()
	case "Bing":
		translator = bing.NewTranslator()
	case "ChatGPT":
		translator = chatgpt.NewTranslator()
	case "DeepL":
		translator = deepl.NewTranslator()
	case "Google":
		translator = google.NewTranslator()
	case "Reverso":
		translator = reverso.NewTranslator()
	}

	return translator
}
