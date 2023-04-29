package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	textURL = "https://api.openai.com/v1/chat/completions"
)

type Translator struct {
	*core.Language
	*core.TTSLock
	core.EngineName
	apiKey string
}

func NewTranslator() *Translator {
	return &Translator{
		Language:   new(core.Language),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("ChatGPT"),
	}
}

func (t *Translator) SetAPIKey(key string) {
	t.apiKey = key
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	if len(t.apiKey) <= 0 {
		return nil, errors.New("Please set API Key in config file for " + t.GetEngineName())
	}

	userData, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{{
			"role": "user",
			"content": fmt.Sprintf(
				"Translate following text from %s to %s\n%s",
				t.GetSrcLang(),
				t.GetDstLang(),
				message,
			),
		}},
		"temperature": 0.7,
	})
	req, _ := http.NewRequest("POST",
		textURL,
		bytes.NewBuffer(userData),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+t.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
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
		data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
