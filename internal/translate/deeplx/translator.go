package deeplx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	textURL = "http://localhost:1188/translate"
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
		EngineName: core.NewEngineName("DeepLX"),
	}
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	userData, _ := json.Marshal(map[string]interface{}{
		"text":        message,
		"source_lang": langCode[t.GetSrcLang()],
		"target_lang": langCode[t.GetDstLang()],
	})
	req, _ := http.NewRequest(http.MethodPost,
		textURL,
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
		return nil, errors.New("translation not found")
	}
	if res.StatusCode != 200 {
		return nil, errors.New(data["message"].(string))
	}

	translation.TEXT = data["data"].(string)

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
