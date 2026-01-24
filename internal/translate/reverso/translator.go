package reverso

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	textURL = "https://api.reverso.net/translate/v1/translation"
	ttsURL  = "https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName=%s?voiceSpeed=80&inputText=%s"
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
		EngineName: core.NewEngineName("Reverso"),
	}
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]any

	if t.GetSrcLang() == t.GetDstLang() {
		return nil, fmt.Errorf("%s doesn't support translation of the same language.\ni.e. %s to %s",
			t.GetEngineName(), t.GetSrcLang(), t.GetDstLang())
	}

	userData, _ := json.Marshal(map[string]any{
		"format": "text",
		"from":   langCode[t.GetSrcLang()],
		"to":     langCode[t.GetDstLang()],
		"input":  message,
		"options": map[string]string{
			"sentenceSplitter":  "true",
			"origin":            "translation.web",
			"contextResults":    "true",
			"languageDetection": "false",
		},
	})
	req, _ := http.NewRequest(http.MethodPost,
		textURL,
		bytes.NewBuffer(userData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", core.UserAgent)
	client := http.DefaultClient
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Request failed, return code: %d", res.StatusCode)
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

	// translation
	for _, line := range data["translation"].([]any) {
		translation.TEXT += line.(string)
	}
	// definition and part of speech
	if data["contextResults"] != nil {
		for _, results := range data["contextResults"].(map[string]any)["results"].([]any) {
			results := results.(map[string]any)
			// definition
			srcExample := results["sourceExamples"].([]any)
			dstExample := results["targetExamples"].([]any)
			if len(srcExample) > 0 && len(dstExample) > 0 {
				for i := 0; i < len(srcExample) && i < len(dstExample); i++ {
					translation.DEF += fmt.Sprintf("- %v\n\t\"%v\"\n", srcExample[i], dstExample[i])
				}
			}
			// part of speech
			if results["partOfSpeech"] == nil {
				translation.POS += fmt.Sprintf("%v\n", results["translation"])
			} else {
				translation.POS += fmt.Sprintf("%v [%v]\n", results["translation"], results["partOfSpeech"])
			}
		}
		translation.DEF = regexp.MustCompile("<(|/)em>").ReplaceAllString(translation.DEF, "")
	}

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	name, ok := voiceName[lang]
	if !ok {
		return errors.New(t.GetEngineName() + " does not support text to speech of " + lang)
	}
	urlStr := fmt.Sprintf(
		ttsURL,
		name,
		base64.StdEncoding.EncodeToString([]byte(message)),
	)
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Add("User-Agent", core.UserAgent)
	client := http.DefaultClient
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Request failed, return code: %d", res.StatusCode)
	}
	return t.Play(res.Body)
}
