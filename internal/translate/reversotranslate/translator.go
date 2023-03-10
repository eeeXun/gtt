package reversotranslate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/eeeXun/gtt/internal/lock"
)

const (
	textURL   = "https://api.reverso.net/translate/v1/translation"
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
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
			"languageDetection": "false",
		},
	})
	req, _ := http.NewRequest("POST",
		textURL,
		bytes.NewBuffer([]byte(userData)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)
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

	// translation
	translation += fmt.Sprintf("%v", data["translation"].([]interface{})[0])
	// definition and part of speech
	if data["contextResults"] != nil {
		for _, results := range data["contextResults"].(map[string]interface{})["results"].([]interface{}) {
			results := results.(map[string]interface{})
			// definition
			srcExample := results["sourceExamples"].([]interface{})
			dstExample := results["targetExamples"].([]interface{})
			if len(srcExample) > 0 && len(dstExample) > 0 {
				for i := 0; i < len(srcExample) && i < len(dstExample); i++ {
					definition += fmt.Sprintf("- %v\n\t\"%v\"\n", srcExample[i], dstExample[i])
				}
			}
			// part of speech
			if results["partOfSpeech"] == nil {
				partOfSpeech += fmt.Sprintf("%v\n", results["translation"])
			} else {
				partOfSpeech += fmt.Sprintf("%v [%v]\n", results["translation"], results["partOfSpeech"])
			}
		}
		definition = regexp.MustCompile("<(|/)em>").ReplaceAllString(definition, "")
	}

	return translation, definition, partOfSpeech, nil
}
