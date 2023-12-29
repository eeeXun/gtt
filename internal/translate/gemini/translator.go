package gemini

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	textURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s"
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
		EngineName: core.NewEngineName("Gemini"),
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

	urlStr := fmt.Sprintf(textURL, t.GetAPIKey())
	userData, _ := json.Marshal(map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": fmt.Sprintf(
							"Translate following text from %s to %s\n%s",
							t.GetSrcLang(),
							t.GetDstLang(),
							message,
						),
					},
				},
			}},
	})
	req, _ := http.NewRequest(http.MethodPost,
		urlStr,
		bytes.NewBuffer(userData),
	)
	req.Header.Add("Content-Type", "application/json")
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
	if data["error"] != nil {
		return nil, errors.New(data["error"].(map[string]interface{})["message"].(string))
	}

	translation.TEXT =
		data["candidates"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["parts"].([]interface{})[0].(map[string]interface{})["text"].(string)

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
