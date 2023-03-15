package reversotranslate

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/eeeXun/gtt/internal/translate/core"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

const (
	textURL   = "https://api.reverso.net/translate/v1/translation"
	ttsURL    = "https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName=%s?voiceSpeed=80&inputText=%s"
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
)

type ReversoTranslate struct {
	*core.Language
	*core.TTSLock
	core.EngineName
}

func NewReversoTranslate() *ReversoTranslate {
	return &ReversoTranslate{
		Language:   core.NewLanguage(),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("ReversoTranslate"),
	}
}

func (t *ReversoTranslate) GetAllLang() []string {
	return lang
}

func (t *ReversoTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data map[string]interface{}

	if t.GetSrcLang() == t.GetDstLang() {
		return "", "", "", errors.New(
			fmt.Sprintf("%s doesn't support translation of the same language.\ni.e. %s to %s",
				t.GetEngineName(), t.GetSrcLang(), t.GetDstLang()))
	}

	userData, _ := json.Marshal(map[string]interface{}{
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

func (t *ReversoTranslate) PlayTTS(lang, message string) error {
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
	req.Header.Add("User-Agent", userAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	decoder, err := mp3.NewDecoder(res.Body)
	if err != nil {
		return err
	}
	otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-readyChan
	player := otoCtx.NewPlayer(decoder)
	player.Play()
	for player.IsPlaying() {
		if t.IsStopped() {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	if err = player.Close(); err != nil {
		return err
	}

	return nil
}
