package reversotranslate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eeeXun/gtt/internal/lock"
)

const (
	textURL = "https://api.reverso.net/translate/v1/translation"
)

type ReversoTranslate struct {
	srcLang    string
	dstLang    string
	EngineName string
	SoundLock  *lock.Lock
}

func (t *ReversoTranslate) GetEngineName() string {
	return t.EngineName
}

func (t *ReversoTranslate) GetAllLang() []string {
	return lang
}

func (t *ReversoTranslate) GetSrcLang() string {
	return t.srcLang
}

func (t *ReversoTranslate) GetDstLang() string {
	return t.dstLang
}

func (t *ReversoTranslate) SetSrcLang(srcLang string) {
	t.srcLang = srcLang
}

func (t *ReversoTranslate) SetDstLang(dstLang string) {
	t.dstLang = dstLang
}

func (t *ReversoTranslate) SwapLang() {
	t.srcLang, t.dstLang = t.dstLang, t.srcLang
}

func (t *ReversoTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data map[string]interface{}

	userData, _ := json.Marshal(map[string]interface{}{
		"format": "text",
		"from":   langCode[t.srcLang],
		"to":     langCode[t.dstLang],
		"input":  message,
		"options": map[string]string{
			"sentenceSplitter":  "true",
			"origin":            "translation.web",
			"contextResults":    "true",
			"languageDetection": "true",
		},
	})
	req, _ := http.NewRequest("POST",
		textURL,
		bytes.NewBuffer([]byte(userData)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", "", err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", "", "", err
	}

	if len(data) <= 0 {
		return "", "", "", errors.New("Translation not found")
	}

	translation += fmt.Sprintf("%v", data["translation"].([]interface{})[0])

	return translation, definition, partOfSpeech, nil
}
