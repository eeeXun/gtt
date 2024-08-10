package libre

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	defaultURL = "https://libretranslate.com/translate"
)

type Translator struct {
	*core.Server
	*core.Language
	*core.TTS
	core.EngineName
}

func NewTranslator() *Translator {
	return &Translator{
		Server:     new(core.Server),
		Language:   new(core.Language),
		TTS:        core.NewTTS(),
		EngineName: core.NewEngineName("Libre"),
	}
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	var textURL string
	if len(t.GetHost()) > 0 {
		textURL = "http://" + t.GetHost() + "/translate"
	} else {
		textURL = defaultURL
	}

	res, err := http.PostForm(textURL,
		url.Values{
			"q":       {message},
			"source":  {langCode[t.GetSrcLang()]},
			"target":  {langCode[t.GetDstLang()]},
			"api_key": {t.GetAPIKey()},
		})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) <= 0 {
		return nil, errors.New("Translation not found")
	}
	if res.StatusCode != 200 {
		return nil, errors.New(data["error"].(string))
	}

	translation.TEXT = data["translatedText"].(string)

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
