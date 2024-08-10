package deepl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	textURL = "https://api-free.deepl.com/v2/translate"
)

type Translator struct {
	*core.Server
	*core.Language
	*core.TTS
	core.EngineName
}

func NewTranslator(name string) *Translator {
	return &Translator{
		Server:     new(core.Server),
		Language:   new(core.Language),
		TTS:        core.NewTTS(),
		EngineName: core.NewEngineName(name),
	}
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) deeplTranslate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	if len(t.GetAPIKey()) <= 0 {
		return nil, errors.New("Please write your API Key in config file for " + t.GetEngineName())
	}

	userData := url.Values{
		"text":        {message},
		"source_lang": {langCode[t.GetSrcLang()]},
		"target_lang": {langCode[t.GetDstLang()]},
	}
	req, _ := http.NewRequest(http.MethodPost,
		textURL,
		strings.NewReader(userData.Encode()),
	)
	req.Header.Add("Authorization", "DeepL-Auth-Key "+t.GetAPIKey())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
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
		return nil, errors.New("translation not found")
	}

	translation.TEXT = data["translations"].([]interface{})[0].(map[string]interface{})["text"].(string)

	return translation, nil
}

func (t *Translator) deeplxTranslate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	if len(t.GetHost()) <= 0 {
		return nil, errors.New("Please write your host in config file for " + t.GetEngineName())
	}

	userData, _ := json.Marshal(map[string]interface{}{
		"text":        message,
		"source_lang": langCode[t.GetSrcLang()],
		"target_lang": langCode[t.GetDstLang()],
	})
	req, _ := http.NewRequest(http.MethodPost,
		"http://"+t.GetHost()+"/translate",
		bytes.NewBuffer(userData),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+t.GetAPIKey())
	res, err := http.DefaultClient.Do(req)
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
		return nil, errors.New(data["message"].(string))
	}

	translation.TEXT = data["data"].(string)

	return translation, nil
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	switch t.GetEngineName() {
	case "DeepLX":
		return t.deeplxTranslate(message)
	default:
		return t.deeplTranslate(message)
	}
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
