package deepl

import (
	"encoding/json"
	"errors"
	"fmt"
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
	*core.APIKey
	*core.Language
	*core.TTSLock
	core.EngineName
}

func NewTranslator() *Translator {
	return &Translator{
		APIKey:     new(core.APIKey),
		Language:   new(core.Language),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("DeepL"),
	}
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	if len(t.GetAPIKey()) <= 0 {
		return nil, errors.New("Please write your API Key in config file for " + t.GetEngineName())
	}

	userData := url.Values{}
	userData.Set("text", message)
	userData.Set("source_lang", langCode[t.GetSrcLang()])
	userData.Set("target_lang", langCode[t.GetDstLang()])
	req, _ := http.NewRequest(http.MethodPost,
		textURL,
		strings.NewReader(userData.Encode()),
	)
	req.Header.Add("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", t.GetAPIKey()))
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

	// fmt.Println("Request: ", req)
	translation.TEXT = data["translations"].([]interface{})[0].(map[string]interface{})["text"].(string)
	// translation.TEXT = fmt.Sprintf("Data: %s\nBody: %s\nRequest: %v", userData.Encode(), body, req)

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
